// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ActivityLog() templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col justify-center items-center gap-10 py-5\"><div class=\"text-center mb-5\"><h2 class=\"text-lg font-bold text-white bg-gray-800 p-3 rounded\">Log Activity</h2></div><form class=\"flex flex-col justify-center items-center gap-5 w-full max-w-md\" hx-post=\"/activity-log\" hx-target=\"#feedback\" hx-swap=\"outerHTML\" hx-indicator=\"#indicator\"><label class=\"input input-bordered flex items-center gap-2 w-full text-white bg-gray-600\">Activity <input type=\"text\" class=\"input input-bordered\" name=\"activity\" required></label> <label class=\"input input-bordered flex items-center gap-2 w-full text-white bg-gray-600\">Start Time <input type=\"datetime-local\" step=\"1\" class=\"input input-bordered\" name=\"start_time\" required></label> <label class=\"input input-bordered flex items-center gap-2 w-full text-white bg-gray-600\">End Time <input type=\"datetime-local\" step=\"1\" class=\"input input-bordered\" name=\"end_time\" required></label> <label class=\"input input-bordered flex items-center gap-2 w-full text-white bg-gray-600\">Comments <input type=\"text\" class=\"input input-bordered\" name=\"comments\" placeholder=\"optional\"></label> <button type=\"submit\" class=\"btn bg-gray-600 text-white hover:bg-gray-700\">Submit activity</button></form><div id=\"indicator\" class=\"htmx-indicator\"><div class=\"flex justify-center items-center w-full\"><span class=\"loading loading-spinner loading-lg text-primary h-20 w-20\"></span></div></div><div id=\"feedback\" class=\"text-center\"></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = template().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
