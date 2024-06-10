// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func template() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link href=\"https://cdn.jsdelivr.net/npm/daisyui@4.10.1/dist/full.min.css\" rel=\"stylesheet\" type=\"text/css\"><script src=\"https://cdn.tailwindcss.com\"></script><script src=\"https://unpkg.com/htmx.org@1.9.11\" integrity=\"sha384-0gxUXCCR8yv9FM2b+U3FDbsKthCI66oH5IA9fHppQq9DDMHuMauqq1ZHBpJxQ0J0\" crossorigin=\"anonymous\"></script><script src=\"https://unpkg.com/htmx.org@1.9.11/dist/ext/response-targets.js\"></script><style>\n            .main-header {\n                font-family: 'Arial', sans-serif;\n                font-size: 24px;\n                font-weight: bold;\n                color: #000;\n                margin: 20px 0;\n            }\n            .nav-buttons a, .nav-buttons button {\n                margin-right: 10px;\n                padding: 10px 15px;\n                border-radius: 5px;\n                text-decoration: none;\n                color: #555; /* Dark gray for better visibility */\n                background: none; /* Removing any background color */\n                transition: color 0.3s; /* Smooth transition for hover effects */\n            }\n            .nav-buttons a:hover, .nav-buttons button:hover {\n                color: #ccc; /* Lighter gray when hovered */\n            }\n            body {\n                background-color: white; /* White background */\n                color: black; /* Black text for better visibility */\n            }\n            .container {\n                max-width: 1200px; /* Ensure container does not exceed screen width */\n            }\n        </style></head><body class=\"bg-white text-gray-900\"><div class=\"container mx-auto p-5\"><div class=\"flex justify-between items-center border-b pb-5 mb-5\"><h1 class=\"main-header\">Symbiont</h1><div class=\"nav-buttons\"><a href=\"/activity-log\">Add Activity</a> <a href=\"/activities-for-day\">Browse Activity</a> <a href=\"/daily-checkin\">Add daily Check-In</a> <a href=\"/daily-checkin-for-day\">Daily Check-In Log</a> <a href=\"/potato-time\">Potato Time</a> <a href=\"/dashboard\">Dashboard</a> <a href=\"/settings\">Settings</a> <button hx-post=\"/logout\">Logout</button></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
