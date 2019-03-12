// Back-end in Go server
// @jeffotoni
// 2019-01-09

package gmail

import "strings"

func TmplDefault(Project, msg string) (tmpl string) {

  Project = strings.Replace(Project, `"`, "", -1)

  tmpl = `<table width="650" border="0" cellpadding="0" cellspacing="0" align="center" style="font-family: Arial,sans-serif; font-size: 14px; color: #2b2722;border:1px solid #E9E9E9;">
  <tr>
    <td colspan="3" style="padding: 5px; text-align: center;"><span style="color: #424242; font-size:11px;">
          Por favor, n&atilde;o responda esse e-mail. Para entrar em contato conosco, 
        <a href="#" style="color: #000; text-decoration: none;"> clique aqui.</a>
        </span></td>
  </tr>
  <tr>
      <td style="padding:30px;text-align:center;background-color:#F8F8F8;" colspan="3"></td>
  </tr>
  <tr>
      <td width="252" colspan="3" style="color:#49B5E2;padding: 30px 30px 10px;font-family: arial;font-size: 20pt;font-weight: bold;text-align:left;">Template Email Default</td>
  </tr>
    
    <tr>
      <td colspan="3" style="color: #424242; font-size: 14px; line-height: 1.5; padding: 0 30px 30px;">
        <p>
        
          Ol&aacute; <span style="font-weight:bold;">Olá Jeff</span>! 
        </p>

        <p>
        ` + msg + `
        </p>
       
      </td>
    </tr>
    
    <tr>
      <td colspan="3" style="font-size:21px; font-weight: bold; border-bottom: 1px solid #e9e9e9;">&nbsp;</td>
    </tr>

  <tr>
    <td colspan="3" style="background-color: #E9E9E9;padding:15px;text-align: center;">
      <a href="#" style="color:#969696;margin:0 5px;display:inline-block;vertical-align:middle;text-decoration:none;">yoursite.com</a>
    </td>
    </tr>
</table>`

  return
}
