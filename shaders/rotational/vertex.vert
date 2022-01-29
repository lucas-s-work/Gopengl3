#version 410
in vec2 vert;
in vec2 verttexcoord;

uniform vec2 dimension;

// Add as needed for now
uniform float rot1angle;
uniform vec2 rot1center;

uniform vec2 trans;

out vec2 fragtexcoord;

void main() {
    fragtexcoord = verttexcoord;

    vec2 pos = vert + trans;

    mat2 rotmat = mat2(
        cos(rot1angle), sin(rot1angle),
        -sin(rot1angle), cos(rot1angle)
    );
    pos = (rotmat * (pos - rot1center)) + rot1center;
    pos.x /= dimension.x * 0.5;
    pos.y /= dimension.y * 0.5;

    gl_Position = vec4(pos - vec2(1.0,1.0), 0., 1.);
}