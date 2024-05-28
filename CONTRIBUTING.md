# Contribution Guide

## Deployment & Initial Setup (Maintainers)

**The following section is intended only for project maintainers & developers**.

1. Install `goreleaser` if it isn't installed

    ```bash
    brew install goreleaser/tap/goreleaser
    ```

2. Initialize the repository for new repositories

    ```bash
    goreleaser init
    ```

3. Test the snapshot without VCS deployment

    ```bash
    goreleaser release --snapshot --clean
    ```

4. Configure the default system's local `gitlab_token` or `github_token` secret

    ```bash
    # mkdir -p ~/.config/goreleaser && touch ~/.config/goreleaser/gitlab_token
    mkdir -p ~/.config/goreleaser && touch ~/.config/goreleaser/github_token
    ```

5. Commit and push to VCS

    ```bash
    git add . && git commit -m "CI - Example" && git push
    ```

6. List `git` tags to get a version that isn't already established

    ```bash
    git tag --list
    ```

7. Using the output from above, increment the version and push a new tag

    ```bash
    git tag -a v0.0.1 -m "Bump: Initial Release" && git push origin v0.0.1
    ```

8. Create and push a new release

    ```bash
    goreleaser release --clean
    ```

9. Update local `Formula` to include the new tag (only applicable for private projects)

    ```bash
    sed -i -e "s/using: GitDownloadStrategy/using: GitDownloadStrategy, tag: \"$(git tag --points-at HEAD)\"/g" ./dist/homebrew/Formula/ethr-cli.rb
    ```

10. Copy the updated `Formula` to system clipboard (only applicable for private projects)

     ```bash
     cat ./dist/homebrew/Formula/ethr-cli.rb | pbcopy
     ```

11. Update the Homebrew Tap's *.rb file with clipboard's contents (only applicable for private projects)

12. Tap the repository using `git+ssh` protocol (only applicable for private repositories)

     ```bash
     brew tap ethr-cli/homebrew-taps git@github.com:x-ethr/homebrew-taps
     ```

13. Update the Cask if already established

     ```bash
     brew update
     ```

14. Install the package (see the [installation](#installing) section)

## Package Publication

_The following section refers to publishing package(s) to https://pkg.go.dev._

- See GO's [*Publishing a Module*](https://go.dev/doc/modules/publishing) for additional details.

1. Establish a [`LICENSE`](https://spdx.org/licenses/) to the project.
2. Ensure dependencies are updated.
    ```bash
    go mod tidy
    ```
3. Sync the working tree's `HEAD` with its remote.
    ```bash
    git add .
    git commit --message "<commit-msg>"
    git push --set-upstream origin main
    ```
4. Assign a tag and push.
    ```bash
    git tag "v$(head VERSION)" && git push origin "v$(head VERSION)"
    ```
5. Make the module available, publicly.
    ```bash
    GOPROXY=proxy.golang.org go list -mutex "github.com/x-ethr/example@v$(head VERSION)"
    ```

Adding the package to `pkg.go.dev` may need to be requested. Navigate to the mirror's expected url, and follow
instructions for requesting the addition.

- Example: https://dev.go.dev/github.com/x-ethr/example

Upon successful request, a message should be displayed:

> _We're still working on “github.com/x-ethr/example”. Check back in a few minutes!_

For any other issues, consult the [official](https://pkg.go.dev/about#adding-a-package) documentation.

### Pre-Commit

The following project makes use of `pre-commit` for local-development `git-hooks`. These hooks are useful
in cases such as preventing secrets from getting pushed into version-control.

See the [`.pre-commit-config.yaml`](.pre-commit-config.yaml) for implementation specifics.

#### Local Setup

1. Install pre-commit from https://pre-commit.com/#install.
2. Auto-update the config to the latest repos' versions by executing `pre-commit autoupdate`.
3. Install with `pre-commit install`.

#### General Command Reference(s)

**Update the configuration's upstreams**

```bash
pre-commit autoupdate
```

**Install `pre-commit` to local instance**

```bash
pre-commit install
```

Tool `godoc` is required to render the documentation, which includes examples.

- See [`doc.go`](./doc.go) for code-specific package documentation.

Installation Steps:

1. Install `godoc`.
    ```bash
    go install golang.org/x/tools/cmd/godoc@latest
    ```
1. Backup shell profile and update `PATH`.
    ```bash
    cp ~/.zshrc ~/.zshrc.bak
    printf "export PATH=\"\${PATH}:%s\"\n" "$(go env --json | jq -r ".GOPATH")/bin" >> ~/.zshrc
    source ~/.zshrc
    ```
1. Start the `godoc` server.
    ```bash
    godoc -http=:6060
    ```
1. Open the webpage.
    ```bash
    open "http://localhost:6060/pkg/"
    ```

## Reference(s)

- [`go` Linking](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications?utm_source=reddit&utm_medium=social&utm_campaign=do-ldflags)
- [`goreleaser`](https://goreleaser.com/install/)
    - [`goreleaser` Environment Variable(s)](https://goreleaser.com/customization/env/)
- [Cobra User Guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)
- [Brew Formula Cookbook](https://github.com/Homebrew/brew/blob/master/docs/Formula-Cookbook.md)
- [New Cask Formula](https://github.com/Homebrew/homebrew-cask)

## Documentation
