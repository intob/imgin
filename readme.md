# imgin
A simple HTTP service to resize images & return output in specified format (JPEG, WEBP or AVIF).

This is inspired by h2non/imaginary. The Dockerfile is also currently based on imaginary.

I had CORS issues with imaginary, so I quickly built this.

There is currently no TLS support because this is an MVP and currently deployed on fly.io (which handles TLS).