package cmd

import (
	"context"
	"fmt"
	"io"
	"sort"
	"regexp"
	"github.com/Masterminds/semver"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/spf13/cobra"
)

// NewCmdCleanup creates a new cobra.Command for the cleanup subcommand.
func NewCmdCleanup(options *[]crane.Option) *cobra.Command {
	keepTags := 3
	repo := ".*"
	tag := ".*"
	dryRun := false
	cmd := &cobra.Command{
		Use:   "cleanup REGISTRY",
		Short: "Cleanup registry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := crane.GetOptions(*options...)

			rrepo, err := regexp.Compile(repo)
			if err != nil {
				return fmt.Errorf("repo-regex %w", err)
			}

			rtag, err := regexp.Compile(tag)
			if err != nil {
				return fmt.Errorf("tag-regex %w", err)
			}

			return filterCatalog(cmd.Context(), cmd.OutOrStdout(), args[0], *rrepo, *rtag, keepTags, dryRun, opts, options)
		},
	}

	cmd.Flags().IntVarP(&keepTags, "keep-tags", "k", 3, "(Optional) Number of last SemVer sorted tags that will not be cleared.")
	cmd.Flags().StringVarP(&repo, "repo-regex", "r", ".*", "(Optional) The regulat expression for the repository selection.")
	cmd.Flags().StringVarP(&tag, "tag-regex", "t", ".*", "(Optional) The regulat expression for the tag selection.")
	cmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "(Optional) Do everything except remove.")

	return cmd
}

func filterCatalog(ctx context.Context, w io.Writer, src string, rrepo regexp.Regexp, 
	               rtag regexp.Regexp, keepTags int, dryRun bool, o crane.Options, options *[]crane.Option) error {
	reg, err := name.NewRegistry(src, o.Name...)

	if err != nil {
		return fmt.Errorf("parsing reg %q: %w", src, err)
	}

	puller, err := remote.NewPuller(o.Remote...)
	if err != nil {
		return err
	}

	catalogger, err := puller.Catalogger(ctx, reg)
	if err != nil {
		return fmt.Errorf("reading tags for %s: %w", reg, err)
	}

	for catalogger.HasNext() {
		repos, err := catalogger.Next(ctx)
		if err != nil {
			return err
		}
		for _, repo := range repos.Repos {
			if rrepo.MatchString(repo) {
				var stags []*semver.Version
				fullrepo, err := name.NewRepository(fmt.Sprintf("%s/%s", src, repo), o.Name...)
				if err != nil {
					return fmt.Errorf("parsing repo %q: %w", src, err)
				}

				lister, err := puller.Lister(ctx, fullrepo)
				if err != nil {
					return fmt.Errorf("reading tags for %s: %w", repo, err)
				}

				for lister.HasNext() {
					tags, err := lister.Next(ctx)

					if err != nil {
						return err
					}
					for _, tag := range tags.Tags {
						if rtag.MatchString(tag) {
							ver, err := semver.NewVersion(tag)
							if err != nil {
								fmt.Fprintln(w, "skip tag", tag, "for repo", fullrepo, "cause", err)
							} else {
								stags = append(stags, ver)
							}
						}
					}
				}

				if len(stags) <= keepTags {
					continue
				}

				sort.Sort(semver.Collection(stags))
				for _, tag := range stags[:(len(stags) - keepTags)] {
					fullname := fmt.Sprintf("%s:%s", fullrepo, tag.Original())
					if !dryRun {
						digest, err := crane.Digest(fullname, *options...)
						if err != nil {
							return fmt.Errorf("geting digest for %s: %w", fullname, err)
						}
						err = crane.Delete(fmt.Sprintf("%s@%s", fullrepo, digest), *options...)
						if err != nil {
							return fmt.Errorf("removing %s: %w", fullname, err)
						}
					}
					fmt.Fprintln(w, "remove", fullname)
				}
			}
		}
	}
	return nil
}
