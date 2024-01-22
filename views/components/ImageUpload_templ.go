// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "prodcat/dto"

func generateImage(img dto.Image) string {
	if img.Body != "" {
		return img.ToString()
	}
	return ""
}

func getAlt(img dto.Image) string {
	return img.Name
}

func ImageUpload(img dto.Image) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := `
		#drop-zone {
			max-width: 450px;
			height: 100%;
			border: 2px dotted blue;
			display: flex;
			justify-content: center;
			align-items: center;
		}

		#drop-zone > img {
			object-fit: cover;
			width: 100%;
			height: 100%;
			display: none;
		}
	`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</style><div id=\"drop-zone\"><img id=\"preview-img\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(generateImage(img)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(getAlt(img)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><p id=\"img-p\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := `Drop file or click to upload`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><input type=\"file\" name=\"uploadimage\" id=\"myfile\" hidden> <input type=\"test\" name=\"test\" id=\"test\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(img.Body))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hidden></div><script type=\"text/javascript\" defer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := `
			var dropZone = document.querySelector('#drop-zone');
			var inputElement = document.querySelector('#myfile');
			var imge = document.querySelector('#preview-img');
			var p = document.querySelector('#img-p')
			var tst = document.querySelector('#test')
			// imge.style = 'display:block';
			// p.style = 'display:none'
			
			console.log(clickFile)
			
			inputElement.addEventListener('change', function (e) {
				const clickFile = this.files[0];

				if (clickFile) {
					imge.style = "display:block;";
					p.style = 'display: none';
					const reader = new FileReader();
					reader.readAsDataURL(clickFile);
					reader.onloadend = function () {
						const result = reader.result;
						let src = this.result;
						imge.src = src;
						imge.alt = clickFile.name
					}
				}
			})

			dropZone.addEventListener('click', () => inputElement.click());
			dropZone.addEventListener('dragover', (e) => {
				e.preventDefault();
			});
			dropZone.addEventListener('drop', (e) => {
				e.preventDefault();
				imge.style = "display:block;";
				let file = e.dataTransfer.files[0];

				const reader = new FileReader();
				reader.readAsDataURL(file);
				reader.onloadend = function () {
					e.preventDefault()
					p.style = 'display: none';
					let src = this.result;
					imge.src = src;
					imge.alt = file.name
				}
			});
	`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
