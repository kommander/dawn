//
// fragment_main
//

void ldexp_2bfc68() {
  float2 res = (2.0f).xx;
}

void fragment_main() {
  ldexp_2bfc68();
}

//
// compute_main
//

void ldexp_2bfc68() {
  float2 res = (2.0f).xx;
}

[numthreads(1, 1, 1)]
void compute_main() {
  ldexp_2bfc68();
}

//
// vertex_main
//
struct VertexOutput {
  float4 pos;
};

struct vertex_main_outputs {
  float4 VertexOutput_pos : SV_Position;
};


void ldexp_2bfc68() {
  float2 res = (2.0f).xx;
}

VertexOutput vertex_main_inner() {
  VertexOutput v = (VertexOutput)0;
  v.pos = (0.0f).xxxx;
  ldexp_2bfc68();
  VertexOutput v_1 = v;
  return v_1;
}

vertex_main_outputs vertex_main() {
  VertexOutput v_2 = vertex_main_inner();
  vertex_main_outputs v_3 = {v_2.pos};
  return v_3;
}

