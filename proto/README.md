# Proto Buffers

Here lives all the protobuffer files with the generated stubs.

## File Structure

The proto files will be stored in the root of this directory under a directory
with domain name. And the generated stubs will be stored under a directory
with the language name and inside that directory the file structure of protos
will be preserved.

For example:

```text
proto
├── go
│   ├── order
│   │   ├── order.pb.go
│   │   └── order_grpc.pb.go
│   ├── payment
│   │   ├── payment.pb.go
│   │   └── payment_grpc.pb.go
│   └── shipping
│       ├── shipping.pb.go
│       └── shipping_grpc.pb.go
├── order
│   └── order.proto
├── payment
│   └── payment.proto
└── shipping
    └── shipping.proto
```

## Generate Stubs

To generate the go stubs, in the root directory of the repo, run the following
command: `make generate-proto`.

## Naming Conventions

- Service name will be named `{domain}Service`. E.g. `PaymentService`.
- The RPC will be named `{action}{domain}`. E.g. `CreatePayment`.
- Request and Response will be named `{action}{domain}{"Request"|"Response"}`.
  E.g. `CreatePaymentRequest`.
