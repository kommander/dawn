#version 310 es

vec3 f(int x) {
  mat3 m = mat3(vec3(1.0f, 2.0f, 3.0f), vec3(4.0f, 5.0f, 6.0f), vec3(7.0f, 8.0f, 9.0f));
  int i = x;
  uint v = min(uint(i), 2u);
  return m[v];
}
layout(local_size_x = 1, local_size_y = 1, local_size_z = 1) in;
void main() {
}
