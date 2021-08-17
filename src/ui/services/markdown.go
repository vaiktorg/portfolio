package services

// import (
// 	"bufio"
// 	"fmt"
// 	"views/template"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/microcosm-cc/bluemonday"
// 	"github.com/vaiktorg/vaiktorg-serv/app/errs"

// 	"github.com/gomarkdown/markdown"
// 	"github.com/gomarkdown/markdown/parser"
// )

// // FullPost contains all data a post will have
// type FullPost struct {
// 	data    PostData
// 	content template.HTML
// }

// // PostData will be used for displaying post previews.
// type PostData struct {
// 	Title         string
// 	Author        string
// 	Description   string
// 	Date          string
// 	ThumbnailPath string
// }

// // RenderPost returns rendered data ready to publish
// func RenderMarkdownBlogPost(path string) (FullPost, PostData, error) {
// 	md, fm := parseMD(path, "---")
// 	if len(md) != 0 && len(fm) != 0 {

// 		prev := PostData{}
// 		post := FullPost{}

// 		// Asserting data coming from frontmatter
// 		if author, ok := fm["author"]; ok {
// 			prev.Author = author
// 		}
// 		if title, ok := fm["title"]; ok {
// 			prev.Title = title
// 		}
// 		if desc, ok := fm["desc"]; ok {
// 			prev.Author = desc
// 		}
// 		if thumb, ok := fm["thumb"]; ok {
// 			if _, err := os.Stat(thumb); os.IsExist(err) {
// 				errs.checkError(err)
// 				prev.ThumbnailPath = thumb
// 			}
// 		}
// 		prev.Date = time.Now().Format("Mon, Jan 2 2006, 15:04:05AM")

// 		// Full Post with Post data struct population

// 		extensions := parser.CommonExtensions | parser.AutoHeadingIDs
// 		parser := parser.NewWithExtensions(extensions)

// 		unsanitizedHTML := markdown.ToHTML([]byte(md), parser, nil)
// 		sanitizedHTML := bluemonday.UGCPolicy().SanitizeBytes(unsanitizedHTML)

// 		post.content = template.HTML(sanitizedHTML)
// 		post.data = prev
// 		return post, prev, nil
// 	}
// 	return FullPost{}, PostData{}, errs.NewError("Couldnt parse files, no data returned")
// }

// // ParseMD returns a map of the key-value pairs found in the front matter.
// func parseMD(path string, delimiter string) (string, map[string]string) {
// 	var markdown string
// 	frontMatterData := makefile(map[string]string)

// 	file, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	scanner := bufio.NewScanner(file)

// 	isFMBlockOpen := false
// 	isFMBlockFinished := false

// 	for scanner.Scan() { //FrontMatter
// 		text := scanner.Text()
// 		if text == delimiter {
// 			if isFMBlockOpen == true {
// 				isFMBlockOpen = false
// 				isFMBlockFinished = true
// 				continue
// 			} else if isFMBlockFinished == false {
// 				isFMBlockOpen = true // Open FM.
// 				continue
// 			}
// 		}
// 		if isFMBlockOpen == true {
// 			splitVals := strings.Split(text, ":")
// 			frontMatterData[splitVals[0]] = strings.TrimSpace(splitVals[1])
// 			continue
// 		}
// 		markdown += fmt.Sprintf(text + "\n")
// 	}
// 	file.Close()

// 	return markdown, frontMatterData
// }
