package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
  strfmt "github.com/go-openapi/strfmt"

  "github.com/go-openapi/errors"
  "github.com/go-openapi/runtime"
  "github.com/go-openapi/validate"
  
  
)



/*DashboardLayout dashboard layout

swagger:model DashboardLayout
*/
type DashboardLayout struct {
  
  
  /* Is Active
 */
Active bool `json:"active,omitempty"`

  
  /* Column Width
 */
ColumnWidth int64 `json:"column_width,omitempty"`

  
  /* Components

Read Only: true
 */
Components [] `json:"components,omitempty"`

  
  /* Id of Dashboard

Read Only: true
 */
DashboardID string `json:"dashboard_id,omitempty"`

  
  /* Unique Id

Read Only: true
 */
ID int64 `json:"id,omitempty"`

  
  /* Type
 */
Type string `json:"type,omitempty"`

  
  /* Width
 */
Width int64 `json:"width,omitempty"`

  
  
  
}


// Validate validates this dashboard layout
func (m *DashboardLayout) Validate(formats strfmt.Registry) error {
  var res []error
  
  

  
  
  
  
  
  
  if err := m.validateComponents(formats); err != nil {
    // prop
    res = append(res, err)
  }
  
  
  
  
  
  
  
  
  
  
  
  

  if len(res) > 0 {
    return errors.CompositeValidationError(res...)
  }
  return nil
}









func (m *DashboardLayout) validateComponents(formats strfmt.Registry) error {
  
  if swag.IsZero(m.Components) { // not required
    return nil
  }
  
  











  return nil
}

















