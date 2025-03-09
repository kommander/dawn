
RWByteAddressBuffer prevent_dce : register(u0);
Texture2D<float4> arg_0 : register(t0, space1);
SamplerState arg_1 : register(s1, space1);
float4 textureSampleBias_a161cf() {
  float4 res = arg_0.SampleBias(arg_1, (1.0f).xx, clamp(1.0f, -16.0f, 15.9899997711181640625f), (int(1)).xx);
  return res;
}

void fragment_main() {
  prevent_dce.Store4(0u, asuint(textureSampleBias_a161cf()));
}

