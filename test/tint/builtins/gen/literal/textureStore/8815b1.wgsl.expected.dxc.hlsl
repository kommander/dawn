//
// fragment_main
//
RWTexture2DArray<int4> arg_0 : register(u0, space1);

void textureStore_8815b1() {
  arg_0[int3((1).xx, int(1u))] = (1).xxxx;
}

void fragment_main() {
  textureStore_8815b1();
  return;
}
//
// compute_main
//
RWTexture2DArray<int4> arg_0 : register(u0, space1);

void textureStore_8815b1() {
  arg_0[int3((1).xx, int(1u))] = (1).xxxx;
}

[numthreads(1, 1, 1)]
void compute_main() {
  textureStore_8815b1();
  return;
}
