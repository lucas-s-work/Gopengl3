#version 410
in vec2 vert;
in vec2 verttexcoord;

uniform vec2 trans;

out vec2 fragtexcoord;

void main() {
    vec2 pos = vert + trans;
    gl_Position = vec4(pos, 0., 1.);
}