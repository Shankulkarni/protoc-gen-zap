# protoc-gen-zap

A protoc plugin that prevents PII from leaking into logs — generates [uber-go/zap](https://github.com/uber-go/zap) `MarshalLogObject` methods for your proto messages with compile-time field redaction.

## The problem

Logging proto structs in Go typically means marshaling to JSON via reflection. This is slow and, more critically, it's easy to accidentally log PII fields like names, emails, or SSNs.

## How it works

protoc-gen-zap generates type-safe `MarshalLogObject` methods for each proto message. Fields tagged with `(zap.redact) = true` are silently dropped from logs — they never appear, not even as empty strings.

**No reflection. No PII leaks. No runtime overhead.**

## Why zap?

Zap's `ObjectMarshaler` interface writes fields directly into a pre-allocated buffer using typed methods (`AddString`, `AddInt64`, etc). Compared to `encoding/json` + reflection:

- ~10x faster marshaling
- Zero allocations in the hot path
- Fields are encoded inline — no intermediate map or struct

## Tag PII fields in your proto

```proto
import "zap/zap.proto";

message User {
    string id             = 1;
    string first_name     = 2 [(zap.redact) = true];  // never logged
    string last_name      = 3 [(zap.redact) = true];  // never logged
    int64  employee_number = 4 [(zap.redact) = true]; // never logged
    string physical_desk  = 6;                        // logged
}
```

## Generated output

```go
func (u *User) MarshalLogObject(enc zapcore.ObjectEncoder) error {
    enc.AddString("Id:", u.Id)
    enc.AddString("PhysicalDesk:", u.PhysicalDesk)
    // first_name, last_name, employee_number — omitted entirely
    return nil
}
```

Redacted fields are gone at compile time — there is no runtime check, no masking, no `"[REDACTED]"` string in your logs.

## Install & run

```bash
go install github.com/ssd-tutorials/protoc-gen-zap

protoc -I ./ \
  --plugin=protoc-gen-zap=protoc-gen-zap \
  --zap_out=:./example \
  ./example/example.proto
```

This generates `example.pb.zap.go` alongside your existing `example.pb.go`.

## Use in logs

```go
logger.Info("user action", zap.Object("user", user))
```

Only non-redacted fields are written. PII fields never reach the log sink.
