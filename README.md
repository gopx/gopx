# What is GoPx and Why?

The existing Golang toolchain is smart enough that it can fetch the dependency packages from any vcs enabled sources like Github or BitBucket, hence it doesn't require any package registry like npm, rubygems. Due its flexible behavior, the package import path has both the source domain name and the package author identifier e.g. `github.com/john/abc`.

But when it comes about API stability and maintainability, the package author and the dependent users might face blocker issues. The package author might need to change his/her identifier or might move the package to any other vcs enabled sources like Github to BitBucket, which causes all the dependent packages to break.

To avoid this issue some gophers use their own domain as the package remote import path and redirects to the actual vcs target server like `john.io/abc` is configured to redirect to `github.com/john/abc`. But this solution doesn't always match to many Go developers. So we must admit that there is a huge gap between package development and package distribution in Golang.

Here the GoPx comes to remove that gap. In GoPx the package author needs to register their package hosting vcs URL to get the stable package remote import path which redirects to the actual registered source, as an example, the author registers his/her package `github.com/john/abc` hosted on Github to GoPx registry as `gopx.io/abc` which is short and uses neither any vcs hosting service name nor the user identifier.

In GoPx every package name is unique like already existing package registry npm or rubygems  etc.

> Stable package registry for Golang

# License

MIT Licensed
