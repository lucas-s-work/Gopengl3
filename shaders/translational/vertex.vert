#version 410
in vec2 vert;
in vec2 verttexcoord;

uniform vec2 trans;

out vec2 fragtexcoord;

void main() {
    fragtexcoord = verttexcoord;
    gl_Position = vec4(vert + trans, 0., 1.);
}