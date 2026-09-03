# Render deployment

Configure these environment variables in Render:

```text
PORT=8080
STATISTICS_API_URL=https://interbank-statistics-api-node.onrender.com
JWT_SECRET=<same-long-secret-used-by-the-other-APIs>
WEB_ORIGIN=https://interbank-matrix-qzaru4ygq-ali-vp.vercel.app
```

`WEB_ORIGIN` must match the Vercel origin exactly: use `https`, no trailing slash. Redeploy after changing it.

The endpoint is available at `https://interbank-matrix-api-go.onrender.com/v1/matrices/qr` and requires `Authorization: Bearer <token>`.
