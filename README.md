# ipsum

## Usage
`ipsum` is command line utility for generating your own "lorem ipsum" from a UTF-8 encoded txt file

Like a good unix-like tool, `ipsum` reads from stdin and outputs to stdout.
```
cat input.txt  | ipsum 
```

You can also specify a file name for either the input and/or the output.
```
ipsum --input script.txt --output output.html
```

By default, ipsum returns 1 paragraph with 5 sentences.
```
ipsum --input script.txt --sentences 4 --paragraph 5
```

For a full list of options
```
ipsum --help
```

## TODO
- Implement an option to remove sentence delimiters from the output
- Implement short flags
- Input is assumed to contain sentence delimiters

## License
Copyright (c) 2016 Daniel Jay Hartman

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.