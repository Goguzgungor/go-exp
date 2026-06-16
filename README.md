# go-exp

Go scriptleri / deneme programları çalıştırmak için hazırlanmış çalışma alanı.

- Go sürümü: 1.24.5
- Modül: `go-exp`

## Yapı

```
go-exp/
├── go.mod              # Go modülü
├── Makefile            # kısayollar
├── cmd/                # her alt klasör çalıştırılabilir bir program
│   └── hello/main.go   # örnek
└── scripts/            # tek dosyalık denemeler (//go:build ignore)
    └── scratch.go
```

## Çalıştırma

İki yaklaşım var:

**1. `cmd/` altında program (önerilen, kalıcı denemeler için)**

Her deneme `cmd/<isim>/main.go` olarak ayrı bir klasöre konur:

```bash
go run ./cmd/hello        # veya: make run CMD=hello
go build -o bin/hello ./cmd/hello
```

Yeni deneme eklemek için:

```bash
mkdir cmd/deneme && cat > cmd/deneme/main.go <<'EOF'
package main
import "fmt"
func main() { fmt.Println("deneme") }
EOF
go run ./cmd/deneme
```

**2. `scripts/` altında tek dosya (hızlı/atılabilir)**

`//go:build ignore` etiketli dosyalar `go build ./...` tarafından yok sayılır,
doğrudan dosya yolu ile çalıştırılır:

```bash
go run scripts/scratch.go
```

## Sık kullanılan komutlar

```bash
make run CMD=hello   # cmd/hello'yu çalıştır
make test            # testleri çalıştır
make fmt             # go fmt
make vet             # go vet
make tidy            # go mod tidy (bağımlılıkları düzenle)
```
