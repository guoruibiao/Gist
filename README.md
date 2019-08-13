# gist

CRUD for gist in console.

## demonstration

0 get your own gist oauth token from [OAuth](https://developer.github.com/v3/gists/)


1 create a gist
```golang
package main

import (
	"github.com/guoruibiao/gist"
	"log"
	"fmt"
)

func main() {
	token := gist.GIST_OAUTH_TOKEN
	gistmanager := gist.NewGistManager(token)
	gist := &gist.Gist{
		Public:true,
		Description:"demonstration for gist-go ",
		Files: map[string]gist.GistFile{
			"filename1": gist.GistFile{
				Content: "file content from filename1",
			},
		},
	}
	success, err := gistmanager.CreateGist(gist)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(success)
}
```

2 delete a gist
```golang
gistmanager := NewGistManager(GIST_OAUTH_TOKEN)
id := "1bd738b1bcb0b1baa2366e3302aa5c5d"
success, err := gistmanager.DeleteGist(id)
if err != nil {
    t.Error(err)
}
t.Log(success)

```

3. update a gist
```golang
gist := &Gist{
    Id:"d04ff9bb10ae9394758f8ae02c78b010",
    Public:true,
    Description:"描述部分信息",
    Files: map[string]GistFile{
        "file1": GistFile{
            Content:"file1的文件正文部分啦啦啦啦",
        },
        "file2": GistFile{
            Content:"document content in file2.🌶😋😝",
        },
    },
}
gistmanager := NewGistManager(GIST_OAUTH_TOKEN)
success, err := gistmanager.UpdateGist(gist)
if err != nil {
    t.Error(err)
}
t.Log(success)

```

4. show a gist
```golang
gistmanager := NewGistManager(GIST_OAUTH_TOKEN)
id := "e021cda97f595181a04151a765084044"
gist, err := gistmanager.ShowGist(id)
if err != nil {
    t.Error(err)
}
t.Log(gist)
```

5. get gist list
```golang
gistmanager := NewGistManager(GIST_OAUTH_TOKEN)
gists, err := gistmanager.ListGists()
if err != nil {
    t.Error(err)
}
t.Log(gists)
t.Log(gists[0])
```

## reference
[gist-REST-API-v3](https://developer.github.com/v3/gists/#list-gist-forks)
