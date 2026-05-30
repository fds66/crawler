package main

import "testing"

func TestGetHeadingFromHTMLBasic(t *testing.T) {

	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "short body",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "longer body",
			inputBody: "<html>  <body>    <h1>Welcome to Boot.dev</h1>    <main>      <p>Learn to code by building real projects.</p>      <p>This is the second paragraph.</p>    </main>  </body></html>",
			expected:  "Welcome to Boot.dev",
		},
		{
			name:      "no h1 but h2",
			inputBody: "<html><body><h2>Test Title</h2></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "neither h1 or h2",
			inputBody: "<html><body><p>Test Title</p></body></html>",
			expected:  "",
		},
		{
			name:      "h1 and h2",
			inputBody: "<html><body>,<h1>This is the first heading</h1><h2>Test Title</h2></body></html>",
			expected:  "This is the first heading",
		},

		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "short body",
			inputBody: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "body with two paragraphs",
			inputBody: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
			<p>Another paragraph.</p>
		</main>
	</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "body with no paragraph",
			inputBody: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<h1>heading</h1>
		</main>
	</body></html>`,
			expected: "",
		},
		{
			name: "empty body",
			inputBody: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			
		</main>
	</body></html>`,
			expected: "",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}
