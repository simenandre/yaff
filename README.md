# YAFF

My name [is YAFF](https://youtu.be/qkBx0gMGuhY?t=19) (pronoused `jef` /dʒɛf/). I am a frontend
framework written in Go, so we don't have to use React (Server Components). 

## Example

You create a file, let's call it `hello-world.yaff`.

```html
---
package main

import (
	"github.com/simenandre/yaff/components"
)

colors := [2]string{"black", "white"}
---

<h1>Hello world</h1>


{#each colors as color}
	<components.HelloWorld color={color} />
{/each}

<style>
h1 {
	padding: 1px:
}
</style>
```

Based on this file, YAFF will generate a file for the markup and Go-code and another
for the styles (including CSS modules).

The templates are structured like this:

```
---
// Component Script (Go-lang)
---
<!-- Component Template (HTML + JS Expressions) -->
```
