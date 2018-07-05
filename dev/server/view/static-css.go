// This file was generated by embed.go. DO NOT EDIT.

package view

var cssResources = map[string][]byte{
	"css/fonts.css": []byte("@font-face {\n\tfont-family: 'Go';\n\tsrc: url('/.static/fonts/GoMedium-Italic.ttf') format('truetype');\n\tfont-weight: 500;\n\tfont-style: italic;\n}\n\n@font-face {\n\tfont-family: 'Go';\n\tsrc: url('/.static/fonts/Go-Italic.ttf') format('truetype');\n\tfont-weight: normal;\n\tfont-style: italic;\n}\n\n@font-face {\n\tfont-family: 'Go';\n\tsrc: url('/.static/fonts/Go-Bold.ttf') format('truetype');\n\tfont-weight: bold;\n\tfont-style: normal;\n}\n\n@font-face {\n\tfont-family: 'Go';\n\tsrc: url('/.static/fonts/GoMedium.ttf') format('truetype');\n\tfont-weight: 500;\n\tfont-style: normal;\n}\n\n@font-face {\n\tfont-family: 'Go';\n\tsrc: url('/.static/fonts/Go-BoldItalic.ttf') format('truetype');\n\tfont-weight: bold;\n\tfont-style: italic;\n}\n\n@font-face {\n\tfont-family: 'Go';\n\tsrc: url('/.static/fonts/GoRegular.ttf') format('truetype');\n\tfont-weight: normal;\n\tfont-style: normal;\n}\n\n@font-face {\n\tfont-family: 'Go Mono';\n\tsrc: url('/.static/fonts/GoMono-Bold.ttf') format('truetype');\n\tfont-weight: bold;\n\tfont-style: normal;\n}\n\n@font-face {\n\tfont-family: 'Go Mono';\n\tsrc: url('/.static/fonts/GoMono.ttf') format('truetype');\n\tfont-weight: normal;\n\tfont-style: normal;\n}\n\n@font-face {\n\tfont-family: 'Go Mono';\n\tsrc: url('/.static/fonts/GoMono-Italic.ttf') format('truetype');\n\tfont-weight: normal;\n\tfont-style: italic;\n}\n\n@font-face {\n\tfont-family: 'Go Mono';\n\tsrc: url('/.static/fonts/GoMono-BoldItalic.ttf') format('truetype');\n\tfont-weight: bold;\n\tfont-style: italic;\n}"),
	"css/main.css": []byte(":root {\n  --link-colour: #05d;\n  --link-hover-colour: #07f;\n  --link-destructive-colour: #d03;\n  --link-destructive-hover-colour: #f06;\n\n  --diagram-channel-colour: #000;\n  --diagram-channel-selected-colour: #09f;\n  --diagram-channel-error-colour: #d03;\n\n  --font-family: Go,'San Francisco','Helvetica Neue',Helvetica,Arial,sans-serif;\n  --font-family-mono: 'Go Mono','Fira Code',Menlo,monospace;\n}\n\nbody {\n\tfont-family: var(--font-family);\n\tfloat: none;\n\tmargin: 0;\n\tdisplay: flex;\n\tflex-flow: column;\n\theight: 100%;\n\tmax-height: 100%;\n}\n\nspan.link {\n\tcolor: var(--link-colour);\n\ttext-decoration: none;\n\tcursor: pointer;\n}\n\nspan.link:hover {\n\tcolor: var(--link-hover-colour);\n\ttext-decoration: underline;\n}\n\nspan.link.destructive {\n    color: var(--link-destructive-colour);\n}\n\nspan.link.destructive:hover {\n    color: var(--link-destructive-hover-colour);\n}\n\na:link, a:visited {\n\tcolor: var(--link-colour);\n\ttext-decoration: none;\n}\n\na:hover {\n\tcolor: var(--link-hover-colour);\n\ttext-decoration: underline;\n}\n\na.destructive:link, a.destructive:visited {\n    color: var(--link-destructive-colour);\n}\n\na.destructive:hover {\n    color: var(--link-destructive-hover-colour);\n}\n\ncode {\n\tfont-family: var(--font-family-mono);\n\tcolor: #066;\n}\n\nform {\n\tfloat: none;\n\tmax-width: 800px;\n\tmargin: 0 auto;\n}\n\ndiv.formfield {\n\tmargin-top: 12px;\n\tmargin-bottom: 12px;\n}\n\ntable {\n\ttable-layout: fixed;\n\tmargin-top: 12px;\n\tmargin-bottom: 12px;\n}\n\ndiv.formfield label {\n\tfloat: left;\n\ttext-align: right;\n\tmargin-right: 15px;\n\twidth: 30%;\n}\n\ninput {\n\tfont-family: var(--font-family-mono);\n\tfont-size: 12pt;\n}\n\ndiv.formfield input[type=text] {\n\twidth: 65%;\n}\n\ndiv.browse-container {\n\tmargin: 0 auto 8px;\n\tmin-width: 800px;\n}\n\nselect {\n\tfont-family: var(--font-family);\n\tfont-size: 12pt;\n}\n\ntextarea {\n\tfont-family: var(--font-family-mono);\n\tfont-size: 12pt;\n}\n\ndiv.head {\n\tpadding: 6px;\n\tflex: 0 1 auto;\n\tborder-bottom-style: solid;\n\tborder-bottom-color: #aaa;\n\tborder-bottom-width: 1px;\n}\n\ndiv.box {\n\tdisplay: flex;\n\tflex-flow: row;\n\tflex: 0 1 auto;\n}\ndiv.container {\n\tflex: 1 1 50%;\n}\n\ndiv#diagram-container {\n\toverflow: scroll;\n}\n\ndiv#panels-container {\n\tpadding: 6px;\n\tdisplay: flex;\n\tflex-flow: column;\n}\n\ndiv.panel {\n\tdisplay: flex;\n\tflex-flow: column;\n\tflex: auto;\n\toverflow: scroll;\n}\n\ndiv.node-panel {\n\tdisplay: flex;\n\tflex-flow: column;\n\tflex: auto;\n}\n\ndiv.hcentre {\n\ttext-align: center;\n}\n\ntable.browse {\n\tfont-family: var(--font-family-mono);\n\tfont-size: 12pt;\n\tmargin-top: 16pt;\n}\n\nfieldset {\n\tmargin: 4px;\n}\n\nfieldset#pathtemplate {\n\tdisplay: none;\n}\n\n.dropdown {\n    position: relative;\n\tdisplay: inline-block;\n\tmargin-left: 2ex;\n}\n\n.dropdown-content {\n    display: none;\n    position: absolute;\n    background-color: #fff;\n    box-shadow: 0px 6px 12px 0px rgba(0,0,0,0.2);\n    padding: 4px 4px;\n    z-index: 1;\n}\n\n.dropdown:hover .dropdown-content {\n    display: block;\n}\n\n.dropdown-content ul {\n\tlist-style-type: none;\n   \tmargin: 0;\n   \tpadding: 0;\n   \toverflow: hidden;\n}\n\n.dropdown-content ul li {\n\twhite-space: nowrap;\n\tmargin: 4px;\n}\n\ndiv.codeedit {\n\tfont-family: var(--font-family-mono);\n\tfont-size: 14px;\n\tflex: auto;\n}\n\ndiv.terminal {\n\tposition: relative;\n\twidth: 100%;\n\theight: 100%;\n}\n\nsvg#diagram {\n\tbackground: #f8f8ff;\n}\n\nsvg#diagram .draggable {\n\tcursor: grab;\n}\n\nsvg#diagram .draggable.dragging {\n\tcursor: grabbing;\n}\n\nsvg#diagram g.textbox text {\n\tfont-family: Go; \n\tfont-size: 16; \n\tuser-select: none; \n\tpointer-events: none;\n\n\talignment-baseline: middle;\n\tdominant-baseline: middle;\n\ttext-anchor: middle;\n}\n\nsvg#diagram g.textbox rect {\n\tfill: #faffee; \n\tstroke: #636e48; \n\tstroke-width: 1;\n}\n\nsvg#diagram g.textbox text::selection {\n    background: none;\n}\n\nsvg#diagram g.node g.textbox rect {\n\tfill: #e0f0ff;\n\tstroke: #45607a;\n\tstroke-width: 1;\n}\n\nsvg#diagram g.node.selected g.textbox rect {\n\tfill: #bee0ff;\n\tstroke-width: 2;\n}\n\nsvg#diagram g.node g.pin circle {\n\tfill: var(--diagram-channel-colour);\n}\n\nsvg#diagram g.node g.pin.selected circle {\n\tfill: var(--diagram-channel-selected-colour);\n}\n\nsvg#diagram g.node g.pin.error circle {\n\tfill: var(--diagram-channel-error-colour);\n}\n\nsvg#diagram g.channel line {\n\tstroke: var(--diagram-channel-colour);\n\tstroke-width: 2;\n}\n\nsvg#diagram g.channel path {\n\tfill: var(--diagram-channel-colour);\n}\n\nsvg#diagram g.channel circle {\n\tfill: var(--diagram-channel-colour);\n}\n\nsvg#diagram g.channel.selected line {\n\tstroke: var(--diagram-channel-selected-colour);\n}\n\nsvg#diagram g.channel.selected path {\n\tfill: var(--diagram-channel-selected-colour);\n}\n\nsvg#diagram g.channel.selected circle {\n\tfill: var(--diagram-channel-selected-colour);\n}\n\nsvg#diagram g.channel.error line {\n\tstroke: var(--diagram-channel-error-colour);\n}\n\nsvg#diagram g.channel.error path {\n\tfill: var(--diagram-channel-error-colour);\n}\n\nsvg#diagram g.channel.error circle {\n\tfill: var(--diagram-channel-error-colour);\n}"),
}