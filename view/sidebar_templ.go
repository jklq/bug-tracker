// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.529
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func sidebar() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"h-screen px-4 pt-3 border-r border-gray-300 flex flex-col items-stretch\" hx-target=\"#page-content\"><span class=\"italic mt-4 block font-semibold select-none\">~bug tracker:</span><div class=\"mt-20\"><a href=\"/app/ticket/new\" hx-boost=\"true\" class=\"p-4 bg-lime-300 whitespace-nowrap rounded  border border-lime-900 duration-100 px-8 hover:bg-lime-400\">+ Add New Ticket</a></div><nav class=\"mt-8 flex-grow\"><ul class=\"flex flex-col gap-1\"><li><a href=\"/app/project\" hx-boost=\"true\" hx-push-url=\"true\" class=\"flex items-center gap-3  p-2 rounded cursor-pointer duration-100 hover:bg-stone-200\"><img class=\"inline h-5 w-5\" src=\"/static/icons/folder-outline.svg\" alt=\"Folder Icon\"> <span class=\"inline-block\">Projects</span></a></li><li><a href=\"/app/ticket\" hx-boost=\"true\" hx-push-url=\"true\" class=\"flex items-center gap-3 hover:bg-stone-200 p-2 rounded cursor-pointer duration-100\"><img class=\"inline h-5 w-5\" src=\"/static/icons/ticket.svg\" alt=\"Folder Icon\"> <span class=\"inline-block\">Assigned Tickets</span></a></li><li><a href=\"/app/project/invitation\" hx-boost=\"true\" hx-push-url=\"true\" class=\"flex items-center gap-3 hover:bg-stone-200 p-2 rounded cursor-pointer duration-100\"><img class=\"inline h-5 w-5\" src=\"/static/icons/mail.svg\" alt=\"Folder Icon\"> <span class=\"inline-block\">Invitations</span></a></li></ul></nav><div class=\"mb-6\"><a href=\"/user/logout\" class=\"p-2 px-4 block flex text-center rounded items-center  border border-lime-900 duration-100 hover:bg-gray-200\"><img src=\"/static/icons/log-out.svg\" class=\"mr-1 h-5 w-5\" alt=\"\">Log out</a></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}