Read a file by path with line numbers; supports offset and line limit (default {{ .DefaultReadLimit }}, max {{ .MaxViewSizeKB }}KB returned file content section); renders images (PNG, JPEG, GIF, WebP); use ls for directories.

LAST-RESORT file reader: prefer tokensave_read (graph lookup) or `rtk read` (shell fallback) first. There is no 'read' tool. Use view only for unindexed files or image rendering.
