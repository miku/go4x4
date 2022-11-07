# Internal package

> internal/ is a special directory name recognised by the go tool which will
> prevent one package from being imported by another unless both share a common
> ancestor.

Example:

> For example, a package /a/b/c/internal/d/e/f can only be imported by code in
> the directory tree rooted at /a/b/c. 

> If your project contains multiple packages you may find you have some exported
> symbols which are intended to be used by other packages in your project, but
> are not intended to be part of your project’s public API.