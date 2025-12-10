# Profiling con pprof

## Iniciar el servidor

```bash
go run main.go
```

## Profiling

### CPU Profile

```bash
# Obtener 30 segundos de CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# En el prompt de pprof:
(pprof) top
(pprof) top10
(pprof) list cpuIntensiveTask
(pprof) web
```

### Memory Profile (Heap)

```bash
go tool pprof http://localhost:6060/debug/pprof/heap

# Ver top allocators
(pprof) top
(pprof) top10 -cum

# Ver gráfico
(pprof) web
```

### Mutex Profile

```bash
go tool pprof http://localhost:6060/debug/pprof/mutex
```

## Benchmark con profiling

```bash
go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=.
go tool pprof cpu.prof
```

## Comparar perfiles

```bash
go tool pprof -base=old.prof new.prof
```

## Visualización

Instalar Graphviz para visualización:
```bash
# macOS
brew install graphviz

# Linux
sudo apt-get install graphviz

# Windows
choco install graphviz
```

Luego usar `web` o `png` en pprof para generar gráficos.

