// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package error_page

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Error404() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"flex flex-col items-center justify-center h-[100vh] gap-4\"><div class=\"items-center justify-center flex flex-col gap-4\"><h1 class=\"text-9xl font-extrabold text-gray-700 tracking-widest\">404</h1><h2 class=\"bg-rose-700 px-2 text-sm rounded rotate-[20deg] absolute\">Resource not found</h2></div><p class=\"text-xs text-center md:text-sm text-gray-400\">The requested resource could not be found but may be available again in the future.</p><a hx-swap=\"transition:true\" href=\"/\" class=\"btn btn-secondary btn-outline\">Home</a></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
