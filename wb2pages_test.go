package wb2pages_test

import (
	"testing"

	wp "github.com/takanoriyanagitani/go-wikibzip2pages"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	t.Run("BasicPages", func(t *testing.T) {
		t.Parallel()

		t.Run("single", func(t *testing.T) {
			t.Parallel()

			const singlePage string = `
  				<page>
  				  	<title>AccessibleComputing</title>
  				  	<ns>0</ns>
  				  	<id>10</id>
  				  	<redirect title="Computer accessibility" />
  				  	<revision>
  				  	  	<id>1219062925</id>
  				  	  	<parentid>1219062840</parentid>
  				  	  	<timestamp>2024-04-15T14:38:04Z</timestamp>
  				  	  	<contributor>
  				  	  		<username>Dummy user name</username>
  				  	  		<id>43603280</id>
  				  	  	</contributor>
  				  	  	<comment>Dummy comment</comment>
  				  	  	<origin>1219062925</origin>
  				  	  	<model>wikitext</model>
  				  	  	<format>text/x-wiki</format>
  				  	  	<text bytes="111" sha1="dummy-sha1" xml:space="preserve">Hello,

  				  	  	World</text>
  				  	  	<sha1>dummy-sha1</sha1>
  				  	</revision>
  				</page>
			`

			pages, e := wp.BasicPages(singlePage)
			if nil != e {
				t.Fatalf("unable to parse: %v\n", e)
			}

			if 1 != len(pages) {
				t.Fatalf("expected single doc. got: %v", len(pages))
			}
		})

	})

	t.Run("BasicPagesBytes", func(t *testing.T) {
		t.Parallel()

		t.Run("single", func(t *testing.T) {
			t.Parallel()

			const singlePage string = `
				<page>
				  	<title>AccessibleComputing</title>
				  	<ns>0</ns>
				  	<id>10</id>
				  	<redirect title="Computer accessibility" />
				  	<revision>
				  	  	<id>1219062925</id>
				  	  	<parentid>1219062840</parentid>
				  	  	<timestamp>2024-04-15T14:38:04Z</timestamp>
				  	  	<contributor>
				  	  		<username>Dummy user name</username>
				  	  		<id>43603280</id>
				  	  	</contributor>
				  	  	<comment>Dummy comment</comment>
				  	  	<origin>1219062925</origin>
				  	  	<model>wikitext</model>
				  	  	<format>text/x-wiki</format>
				  	  	<text bytes="111" sha1="dummy-sha1" xml:space="preserve">Hello,

				  	  	World</text>
				  	  	<sha1>dummy-sha1</sha1>
				  	</revision>
				</page>
			`

			// convert the string to a byte slice
			pageBytes := []byte(singlePage)

			pages, e := wp.BasicPagesBytes(pageBytes)
			if e != nil {
				t.Fatalf("unable to parse bytes: %v\n", e)
			}

			if len(pages) != 1 {
				t.Fatalf("expected single doc from bytes. got: %v", len(pages))
			}
			if pages[0].Title != "AccessibleComputing" {
				t.Fatalf("unexpected title from bytes: %q", pages[0].Title)
			}
		})

		t.Run("multi", func(t *testing.T) {
			t.Parallel()

			const multiPages string = `
 				<page>
 				  	<title>FirstPage</title>
 				  	<ns>0</ns>
 				  	<id>1</id>
 				  	<revision>
 				  		<id>101</id>
 				  		<timestamp>2024-01-01T00:00:00Z</timestamp>
 				  		<text xml:space="preserve">Hello first page</text>
 				  	</revision>
 				</page>
 				<page>
 				  	<title>SecondPage</title>
 				  	<ns>0</ns>
 				  	<id>2</id>
 				  	<revision>
 				  		<id>102</id>
 				  		<timestamp>2024-01-02T00:00:00Z</timestamp>
 				  		<text xml:space="preserve">Hello second page</text>
 				  	</revision>
 				</page>
 			`

			pages, err := wp.BasicPagesBytes([]byte(multiPages))
			if err != nil {
				t.Fatalf("unable to parse bytes: %v\n", err)
			}

			if len(pages) != 2 {
				t.Fatalf("expected 2 pages, got: %d", len(pages))
			}

			if pages[0].Title != "FirstPage" || pages[1].Title != "SecondPage" {
				t.Fatalf("unexpected titles: %q, %q", pages[0].Title, pages[1].Title)
			}
		})

	})
}
