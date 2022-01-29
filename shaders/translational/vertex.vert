#version 410
in vec2 vert;
in vec2 verttexcoord;

uniform vec2 trans;

out vec2 fragtexcoord;

void main() {
    fragtexcoord = verttexcoord;
    vec2 pos = vert + trans;

    pos.x /= dimensions.x * 0.5
    pos.y /= dimensions.y * 0.5

    gl_Position = vec4(pos - vec2(1.0,1.0), 0., 1.);
}