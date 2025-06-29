package go_shopware_admin_sdk

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/shyim/go-version"
)

type ExtensionManagerService ClientService

func (e ExtensionManagerService) Refresh(ctx ApiContext) (*http.Response, error) {
	r, err := e.Client.NewRequest(ctx, http.MethodPost, "/api/_action/extension/refresh", nil)

	if err != nil {
		return nil, fmt.Errorf("cannot refresh extension manager %w", err)
	}

	return e.Client.BareDo(ctx.Context, r)
}

func (e ExtensionManagerService) ListAvailableExtensions(ctx ApiContext) (ExtensionList, *http.Response, error) {
	r, err := e.Client.NewRequest(ctx, http.MethodGet, "/api/_action/extension/installed", nil)

	if err != nil {
		return nil, nil, fmt.Errorf("cannot list installed extensions %w", err)
	}

	var extensions ExtensionList
	resp, err := e.Client.Do(ctx.Context, r, &extensions)

	if err != nil {
		return nil, nil, err
	}

	return extensions, resp, err
}

func (e ExtensionManagerService) lifecycleUpdate(typeName string, ctx ApiContext, httpUrl, httpMethod string) (*http.Response, error) {
	r, err := e.Client.NewRequest(ctx, httpMethod, httpUrl, nil)

	if err != nil {
		return nil, fmt.Errorf("cannot %s %w", typeName, err)
	}

	return e.Client.BareDo(ctx.Context, r)
}

func (e ExtensionManagerService) InstallExtension(ctx ApiContext, extType, name string) (*http.Response, error) {
	return e.lifecycleUpdate("InstallExtension", ctx, fmt.Sprintf("/api/_action/extension/install/%s/%s", extType, name), http.MethodPost)
}

func (e ExtensionManagerService) UninstallExtension(ctx ApiContext, extType, name string) (*http.Response, error) {
	return e.lifecycleUpdate("UninstallExtension", ctx, fmt.Sprintf("/api/_action/extension/uninstall/%s/%s", extType, name), http.MethodPost)
}

func (e ExtensionManagerService) UpdateExtension(ctx ApiContext, extType, name string) (*http.Response, error) {
	return e.lifecycleUpdate("UpdateExtension", ctx, fmt.Sprintf("/api/_action/extension/update/%s/%s", extType, name), http.MethodPost)
}

func (e ExtensionManagerService) DownloadExtension(ctx ApiContext, name string) (*http.Response, error) {
	return e.lifecycleUpdate("DownloadExtension", ctx, fmt.Sprintf("/api/_action/extension/download/%s", name), http.MethodPost)
}

func (e ExtensionManagerService) ActivateExtension(ctx ApiContext, extType, name string) (*http.Response, error) {
	return e.lifecycleUpdate("ActivateExtension", ctx, fmt.Sprintf("/api/_action/extension/activate/%s/%s", extType, name), http.MethodPut)
}

func (e ExtensionManagerService) DeactivateExtension(ctx ApiContext, extType, name string) (*http.Response, error) {
	return e.lifecycleUpdate("ActivateExtension", ctx, fmt.Sprintf("/api/_action/extension/deactivate/%s/%s", extType, name), http.MethodPut)
}

func (e ExtensionManagerService) RemoveExtension(ctx ApiContext, extType, name string) (*http.Response, error) {
	// Since 6.6.10.2 is it POST instead of DELETE
	if version.MustConstraints(version.NewConstraint(">=6.6.10.2")).Check(e.Client.ShopwareVersion) {
		return e.lifecycleUpdate("RemoveExtension", ctx, fmt.Sprintf("/api/_action/extension/remove/%s/%s", extType, name), http.MethodPost)
	}

	return e.lifecycleUpdate("RemoveExtension", ctx, fmt.Sprintf("/api/_action/extension/remove/%s/%s", extType, name), http.MethodDelete)
}

func (e ExtensionManagerService) UploadExtension(ctx ApiContext, extensionZip io.Reader) (*http.Response, error) {
	var buf bytes.Buffer
	parts := multipart.NewWriter(&buf)
	mimeHeader := textproto.MIMEHeader{}
	mimeHeader.Set("Content-Disposition", `form-data; name="file"; filename="extension.zip"`)
	mimeHeader.Set("Content-Type", "application/zip")

	part, err := parts.CreatePart(mimeHeader)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(part, extensionZip); err != nil {
		return nil, err
	}
	if err := parts.Close(); err != nil {
		return nil, err
	}

	var body io.Reader = &buf

	r, err := e.Client.NewRawRequest(ctx, http.MethodPost, "/api/_action/extension/upload", body)

	if err != nil {
		return nil, fmt.Errorf("cannot upload extension %w", err)
	}

	r.Header.Set("Content-Type", parts.FormDataContentType())

	return e.Client.BareDo(ctx.Context, r)
}

func (e ExtensionManagerService) UploadExtensionUpdateToCloud(ctx ApiContext, extensionName string, extensionZip io.Reader) (*http.Response, error) {
	var buf bytes.Buffer
	parts := multipart.NewWriter(&buf)

	if writer, err := parts.CreateFormField("media"); err != nil {
		return nil, err
	} else {
		_, err := writer.Write([]byte(extensionName))
		if err != nil {
			return nil, err
		}
	}

	mimeHeader := textproto.MIMEHeader{}
	mimeHeader.Set("Content-Disposition", `form-data; name="file"; filename="extension.zip"`)
	mimeHeader.Set("Content-Type", "application/zip")

	part, err := parts.CreatePart(mimeHeader)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(part, extensionZip); err != nil {
		return nil, err
	}
	if err := parts.Close(); err != nil {
		return nil, err
	}

	var body io.Reader = &buf

	r, err := e.Client.NewRawRequest(ctx, http.MethodPost, "/api/_action/extension/update-private", body)

	if err != nil {
		return nil, fmt.Errorf("cannot upload extension update to cloud %w", err)
	}

	r.Header.Set("Content-Type", parts.FormDataContentType())

	return e.Client.BareDo(ctx.Context, r)
}
