#version 310 es


struct S_std140 {
  int before;
  uint tint_pad_0;
  uint tint_pad_1;
  uint tint_pad_2;
  vec3 m_col0;
  uint tint_pad_3;
  vec3 m_col1;
  uint tint_pad_4;
  vec3 m_col2;
  uint tint_pad_5;
  int after;
  uint tint_pad_6;
  uint tint_pad_7;
  uint tint_pad_8;
  uint tint_pad_9;
  uint tint_pad_10;
  uint tint_pad_11;
  uint tint_pad_12;
  uint tint_pad_13;
  uint tint_pad_14;
  uint tint_pad_15;
  uint tint_pad_16;
  uint tint_pad_17;
  uint tint_pad_18;
  uint tint_pad_19;
  uint tint_pad_20;
};

struct S {
  int before;
  mat3 m;
  int after;
};

layout(binding = 0, std140)
uniform u_block_std140_1_ubo {
  S_std140 inner[4];
} v;
S p[4] = S[4](S(0, mat3(vec3(0.0f), vec3(0.0f), vec3(0.0f)), 0), S(0, mat3(vec3(0.0f), vec3(0.0f), vec3(0.0f)), 0), S(0, mat3(vec3(0.0f), vec3(0.0f), vec3(0.0f)), 0), S(0, mat3(vec3(0.0f), vec3(0.0f), vec3(0.0f)), 0));
S tint_convert_S(S_std140 tint_input) {
  return S(tint_input.before, mat3(tint_input.m_col0, tint_input.m_col1, tint_input.m_col2), tint_input.after);
}
layout(local_size_x = 1, local_size_y = 1, local_size_z = 1) in;
void main() {
  S_std140 v_1[4] = v.inner;
  S v_2[4] = S[4](S(0, mat3(vec3(0.0f), vec3(0.0f), vec3(0.0f)), 0), S(0, mat3(vec3(0.0f), vec3(0.0f), vec3(0.0f)), 0), S(0, mat3(vec3(0.0f), vec3(0.0f), vec3(0.0f)), 0), S(0, mat3(vec3(0.0f), vec3(0.0f), vec3(0.0f)), 0));
  {
    uint v_3 = 0u;
    v_3 = 0u;
    while(true) {
      uint v_4 = v_3;
      if ((v_4 >= 4u)) {
        break;
      }
      v_2[v_4] = tint_convert_S(v_1[v_4]);
      {
        v_3 = (v_4 + 1u);
      }
      continue;
    }
  }
  p = v_2;
  p[1u] = tint_convert_S(v.inner[2u]);
  p[3u].m = mat3(v.inner[2u].m_col0, v.inner[2u].m_col1, v.inner[2u].m_col2);
  p[1u].m[0u] = v.inner[0u].m_col1.zxy;
}
