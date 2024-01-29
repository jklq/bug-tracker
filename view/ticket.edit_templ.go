// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.529
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/jklq/bug-tracker/db"

func TicketEditView(template templ.Component, ticket db.Ticket) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"border-b border-gray-300 p-4 px-8\"><div class=\"flex max-w-screen-lg items-stretch gap-3\"><div><a href=\"view\" hx-get=\"view\" hx-push-url=\"true\" class=\"p-1 h-full flex items-center block rounded hover:bg-gray-200 border border-lime-900 duration-100 px-3\">&lt;</a></div><div class=\"flex-grow\"><h1 class=\"text-xl font-bold\">Edit \"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(ticket.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/ticket.edit.templ`, Line: 19, Col: 55}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"</h1><p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(ticket.Description.String)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/ticket.edit.templ`, Line: 20, Col: 35}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div></div><div class=\"flex max-w-screen-lg\"><div class=\"flex-grow\"><form hx-push-url=\"true\" action=\"edit\" class=\"p-4 px-8\" hx-target-error=\"#project-edit-error\" method=\"post\"><!-- Ticket Title --><div class=\"mb-2\"><label for=\"ticket-name-input\" class=\"block\">Ticket Title:</label> <input id=\"ticket-name-input\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(ticket.Title))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full border border-gray-700 p-1 text-lg\" type=\"text\" name=\"title\"></div><!-- Ticket Description --><div class=\"mb-2\"><label for=\"ticket-description-input\" class=\"block\">Ticket Description:</label> <textarea id=\"ticket-description-input\" class=\"w-full border border-gray-700 p-1 text-lg\" type=\"text\" name=\"description\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(ticket.Description.String)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/ticket.edit.templ`, Line: 46, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</textarea></div><!-- Ticket Status --><div class=\"mb-2\"><label for=\"ticket-status-input\" class=\"block\">Ticket Status:</label> <select id=\"ticket-status-input\" class=\"w-full border border-gray-700 p-1 text-lg\" name=\"status\"><option value=\"1\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if ticket.Status == 1 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Open</option> <option value=\"2\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if ticket.Status == 2 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">In Progress</option> <option value=\"0\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if ticket.Status == 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Closed</option></select></div><!-- Ticket Priority --><div class=\"mb-2\"><label for=\"ticket-priority-input\" class=\"block\">Ticket Priority:</label> <select id=\"ticket-priority-input\" class=\"w-full border border-gray-700 p-1 text-lg\" name=\"priority\"><option value=\"1\">Low</option> <option value=\"2\">Medium</option> <option value=\"3\">High</option></select></div><!-- Submit Button --><button type=\"submit\" class=\"p-2 border border-black hover:bg-gray-300 duration-100 mt-3\">Save changes</button><!-- Error Display --><div class=\"text-red-600 inline-block p-2\" id=\"project-edit-error\"></div></form></div><div class=\"p-4 \"><h2 class=\"text-xl text-bold mb-2\">Other actions</h2><div class=\"p-3 border border-red-800 rounded\"><form action=\"post\" method=\"post\" hx-post=\"delete\" hx-push-url=\"true\" hx-push-url=\"true\" method=\"post\"><input type=\"submit\" class=\"p-1 py-2 bg-red-600 cursor-pointer rounded hover:bg-red-700 text-white duration-100 px-8\" value=\"Delete project\"></form></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = template.Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}