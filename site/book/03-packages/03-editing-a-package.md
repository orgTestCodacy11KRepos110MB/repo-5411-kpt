kpt does not maintain any state on your local machine outside of the directory where you fetched the
package. Making changes to the package is accomplished by manipulating the local filesystem. At the
lowest-level, _editing_ a package is simply a process that either:

- Changes the resources within that package. Examples:
  - Authoring new a Deployment resource
  - Customizing an existing Deployment resource
  - Modifying the Kptfile or other meta resources
- Changes the package hierarchy, also called _package composition_. Examples:
  - Adding a subpackage.
  - Create a new dependent subpackage.

At the end of the day, editing a package will result in a Git commit that fully specifies
the package. This process can be manual or automated depending on your use case.

We will cover package composition later in this chapter. For now, let's focus on editing resources
_within_ a package.

## Manual edits

As mentioned earlier, you can manually edit or author KRM resources using your favorite editor.
Since every KRM resource has a known schema, you can take advantage of tooling that assist in
authoring and validating resource configuration. For example, [Cloud Code] extensions for VS Code
and IntelliJ provide IDE features such as auto-completion, inline documentation, linting, and snippets.

For example, if you have VS Code installed, try modifying the resources in the `wordpress` package:

```shell
$ code wordress
```

## Automation

Often times you want to automate repetitive or complex operations. Having standardized on KRM for
all resources in a package (including meta resources) allows us to easily develop automation in
different toolchains, languages, and at levels of abstraction.

For example, setting a label on all the resources in the `wordpress` package can be done
using the following function:

```shell
cd wordpress
$ kpt fn eval --image gcr.io/kpt-fn/set-label:v0.1 -- env=dev
```

[Chapter 4] discusses different ways of running functions in detail.

[cloud code]: https://cloud.google.com/code
[chapter 4]: /book/04-using-functions/