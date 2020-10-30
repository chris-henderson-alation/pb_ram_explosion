use tonic_build;

fn main() {
    pb();
}

fn pb() {
    tonic_build::compile_protos("./example.proto").unwrap();
}
