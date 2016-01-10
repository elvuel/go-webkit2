package webkit2

// #include <webkit2/webkit2.h>
import "C"
import "unsafe"
import "github.com/gotk3/gotk3/glib"

type Settings struct {
	*glib.Object
	settings *C.WebKitSettings
}

// newSettings creates a new Settings with default values.
//
// See also: webkit_settings_new at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-new.
func newSettings(settings *C.WebKitSettings) *Settings {
	return &Settings{&glib.Object{glib.ToGObject(unsafe.Pointer(settings))}, settings}
}

// GetAutoLoadImages returns the "auto-load-images" property.
//
// See also: webkit_settings_get_auto_load_images at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-get-auto-load-images
func (s *Settings) GetAutoLoadImages() bool {
	return gobool(C.webkit_settings_get_auto_load_images(s.settings))
}

// SetAutoLoadImages sets the "auto-load-images" property.
//
// See also: webkit_settings_get_auto_load_images at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-set-auto-load-images
func (s *Settings) SetAutoLoadImages(autoLoad bool) {
	C.webkit_settings_set_auto_load_images(s.settings, gboolean(autoLoad))
}

// SetUserAgentWithApplicationDetails sets the "user-agent" property by
// appending the application details to the default user agent.
//
// See also: webkit_settings_set_user_agent_with_application_details at
// http://webkitgtk.org/reference/webkit2gtk/unstable/WebKitSettings.html#webkit-settings-set-user-agent-with-application-details
func (s *Settings) SetUserAgentWithApplicationDetails(appName, appVersion string) {
	C.webkit_settings_set_user_agent_with_application_details(s.settings, (*C.gchar)(C.CString(appName)), (*C.gchar)(C.CString(appVersion)))
}

// SetUserAgent
//
// See also: webkit_settings_set_user_agent at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-set-user-agent
func (s *Settings) SetUserAgent(ua string) {
	C.webkit_settings_set_user_agent(s.settings, (*C.gchar)(C.CString(ua)))
}

// BELOWS COULD ALSO USE #SetProperty, add for T
// T
func (s *Settings) SetEnableOfflineWebApplicationCache(v bool) {
	C.webkit_settings_set_enable_offline_web_application_cache(s.settings, gboolean(v))
}

// T
func (s *Settings) SetJavascriptCanOpenWindowsAutomatically(v bool) {
	C.webkit_settings_set_javascript_can_open_windows_automatically(s.settings, gboolean(v))
}

// T
func (s *Settings) SetEnablePrivateBrowsing(v bool) {
	C.webkit_settings_set_enable_private_browsing(s.settings, gboolean(v))
}

// T
func (s *Settings) SetEnableDeveloperExtras(v bool) {
	C.webkit_settings_set_enable_developer_extras(s.settings, gboolean(v))
}

// T
func (s *Settings) SetEnableWebaudio(v bool) {
	C.webkit_settings_set_enable_webaudio(s.settings, gboolean(v))
}

// T
func (s *Settings) SetEnablePageCache(v bool) {
	C.webkit_settings_set_enable_page_cache(s.settings, gboolean(v))
}
