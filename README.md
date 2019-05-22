# Homomorphine-Go
Homomorphine-Go is a Golang interface for [Homomorphine](https://github.com/caboom/homomorphine-go) library - universal backend for various Homomorphic encryption backends and algorithms.

# Installation
You have to previously install [Homomorphine](https://github.com/caboom/homomorphine-go) and relevant homomorphic libraries you intend to use (SEAL, HELib, etc.)

You can install Homomorphine-Go  package with:
```
make all
make install
```

*Note* Homomorphine-Go requires cgo as it is using Homomorphine's clang backend.

# How to use it?

At the moment it's best to check the tests, however, more examples are coming in near future.

# Warning 

Homomorphine-Go is currently not ready for production use.
