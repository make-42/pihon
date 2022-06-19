include <BOSL/constants.scad>
use <BOSL/masks.scad>

// Start config
// Pi Board
pi_board_size = [65,30,13]; //mm
pi_board_screw_holes_center_from_board_edge_distance = 3.5; //mm
pi_board_screw_holes_diameter = 2.75; //mm
pi_board_padding = 2; //mm

// Ports
pi_board_usb_offset_from_edge = 6; //mm
pi_board_usb_space_between_ports = 12.6; //mm
pi_board_hdmi_offset_from_edge = 46.85-1; //mm
pi_board_usb_size = [9,10,5]; //mm
pi_board_hdmi_size = [11.5,10,6]; //mm
pi_board_port_rounding_radius = 1; //mm

// Screens
screen_pcb_size = [24.7,27,13.3]; //mm (+- 0.2) 13.3mm are pin headers.
screen_pcb_screw_holes_center_from_board_edge_distance = 2.1; //mm
screen_pcb_screw_holes_diameter=2.8; //mm
active_screen_area = [24.7,13.86,3.5]; //mm
active_screen_area_offset_from_corner = [0,5.08]; //mm

// Push buttons
push_button_size = [6,6,6.5]; //mm
push_button_offset_from_corners = 5; //mm
push_button_support_thickness = 2;//mm
push_button_distance_from_case_bottom = 12; //mm
// Compliant thumbs
push_button_compliant_thumb_tolerence = 1; //mm

// Misc
walls_thickness = 1; //mm
z_padding = 5; //mm
rounding_radius=0.5; //mm

resolution = 80;
// End config

// Modules
module sector(radius, angles, fn = resolution) {
    r = radius / cos(180 / fn);
    step = -360 / fn;

    points = concat([[0, 0]],
        [for(a = [angles[0] : step : angles[1] - 360]) 
            [r * cos(a), r * sin(a)]
        ],
        [[r * cos(angles[1]), r * sin(angles[1])]]
    );

    difference() {
        circle(radius, $fn = fn);
        polygon(points);
    }
}

module arc(radius, angles, width = 1, fn = resolution) {
    difference() {
        sector(radius + width, angles, fn);
        sector(radius, angles, fn);
    }
}

// Modeling
difference(){
// Body
cube(size=[pi_board_size.x+2*pi_board_padding+walls_thickness*2,pi_board_size.y+2*pi_board_padding+2*walls_thickness,pi_board_size.z+screen_pcb_size.z+z_padding+walls_thickness]);
   // Fillet Edge
    fillet_mask_z(l=pi_board_size.z+screen_pcb_size.z+z_padding+walls_thickness, r=rounding_radius, align=V_UP, $fn=resolution);
    
   translate([pi_board_size.x+2*pi_board_padding+walls_thickness*2,pi_board_size.y+2*pi_board_padding+2*walls_thickness,0])
    fillet_mask_z(l=pi_board_size.z+screen_pcb_size.z+z_padding+walls_thickness, r=rounding_radius, align=V_UP, $fn=resolution);
    
   translate([pi_board_size.x+2*pi_board_padding+walls_thickness*2,0,0]) fillet_mask_z(l=pi_board_size.z+screen_pcb_size.z+z_padding+walls_thickness, r=rounding_radius, align=V_UP, $fn=resolution);
    
   translate([0,pi_board_size.y+2*pi_board_padding+2*walls_thickness,0]) fillet_mask_z(l=pi_board_size.z+screen_pcb_size.z+z_padding+walls_thickness, r=rounding_radius, align=V_UP, $fn=resolution);
    
translate([walls_thickness,walls_thickness,walls_thickness])
cube(size=[pi_board_size.x+2*pi_board_padding,pi_board_size.y+2*pi_board_padding,pi_board_size.z+screen_pcb_size.z+z_padding]);
// pi board screws
    //1
    translate([walls_thickness,pi_board_size.y+pi_board_padding*2+walls_thickness,0])
    translate([pi_board_padding,-pi_board_padding])
    translate([pi_board_screw_holes_center_from_board_edge_distance,-pi_board_screw_holes_center_from_board_edge_distance])
    cylinder(5,pi_board_screw_holes_diameter/2,pi_board_screw_holes_diameter/2,$fn=resolution);
    //2
    translate([walls_thickness,pi_board_size.y+pi_board_padding*2+walls_thickness,0])
    translate([pi_board_padding,-pi_board_padding])
    translate([pi_board_size.x-pi_board_screw_holes_center_from_board_edge_distance,-pi_board_screw_holes_center_from_board_edge_distance,0])
    cylinder(5,pi_board_screw_holes_diameter/2,pi_board_screw_holes_diameter/2,$fn=resolution);
    //3
    translate([walls_thickness,walls_thickness,0])
    translate([pi_board_size.x+pi_board_padding,pi_board_padding,0])
    translate([-pi_board_screw_holes_center_from_board_edge_distance,pi_board_screw_holes_center_from_board_edge_distance,0])
    cylinder(5,pi_board_screw_holes_diameter/2,pi_board_screw_holes_diameter/2,$fn=resolution);
    //4
    translate([walls_thickness,walls_thickness,0])
    translate([pi_board_padding,pi_board_padding,0])
    translate([pi_board_screw_holes_center_from_board_edge_distance,pi_board_screw_holes_center_from_board_edge_distance,0])
    cylinder(5,pi_board_screw_holes_diameter/2,pi_board_screw_holes_diameter/2,$fn=resolution);
    
    // ports
    // USB 1
    translate([walls_thickness,pi_board_size.y+pi_board_padding*2+walls_thickness,walls_thickness])
    translate([pi_board_padding+pi_board_usb_offset_from_edge,-pi_board_usb_size.y+walls_thickness,0])
    difference(){
        cube(pi_board_usb_size);
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([pi_board_usb_size.x,0,0])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([pi_board_usb_size.x,0,pi_board_usb_size.z])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([0,0,pi_board_usb_size.z])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
    }
        

    
    // USB 2
translate([walls_thickness,pi_board_size.y+pi_board_padding*2+walls_thickness,walls_thickness])
translate([pi_board_usb_space_between_ports,0,0])
    translate([pi_board_padding+pi_board_usb_offset_from_edge,-pi_board_usb_size.y+walls_thickness,0])
    difference(){
        cube(pi_board_usb_size);
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([pi_board_usb_size.x,0,0])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([pi_board_usb_size.x,0,pi_board_usb_size.z])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([0,0,pi_board_usb_size.z])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
    }
    // HDMI 1
translate([walls_thickness,pi_board_size.y+pi_board_padding*2+walls_thickness,walls_thickness])
    translate([pi_board_padding+pi_board_hdmi_offset_from_edge,-pi_board_usb_size.y+walls_thickness,0])
    difference(){
        cube(pi_board_hdmi_size);
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([pi_board_hdmi_size.x,0,0])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([pi_board_hdmi_size.x,0,pi_board_hdmi_size.z])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
        translate([0,0,pi_board_hdmi_size.z])
        fillet_mask_y(l=100, r=pi_board_port_rounding_radius, align=V_BACK,$fn=resolution);
    }
    
    // Button compliant thumbs
    //1
    translate([0,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x-push_button_offset_from_corners,walls_thickness])
    translate([0,0,0])
    cube(size=[walls_thickness,push_button_compliant_thumb_tolerence,push_button_distance_from_case_bottom+push_button_size.x-(push_button_size.x/2+push_button_compliant_thumb_tolerence/2)]);
    translate([0,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x-push_button_offset_from_corners,walls_thickness])
    translate([0,push_button_size.x,0])
    cube(size=[walls_thickness,push_button_compliant_thumb_tolerence,push_button_distance_from_case_bottom+push_button_size.x-(push_button_size.x/2+push_button_compliant_thumb_tolerence/2)]);
    
    translate([0,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x-push_button_offset_from_corners,walls_thickness])
    translate([0,push_button_size.x/2+push_button_compliant_thumb_tolerence/2,push_button_distance_from_case_bottom+push_button_size.x-(push_button_size.x/2+push_button_compliant_thumb_tolerence/2)])
    rotate([90,0,90])
    linear_extrude(walls_thickness) arc(push_button_size.x/2-push_button_compliant_thumb_tolerence/2, [0,180], push_button_compliant_thumb_tolerence);
    //2
    translate([pi_board_size.x+2*pi_board_padding+walls_thickness,0,0])
    translate([0,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x-push_button_offset_from_corners,walls_thickness])
    translate([0,0,0])
    cube(size=[walls_thickness,push_button_compliant_thumb_tolerence,push_button_distance_from_case_bottom+push_button_size.x-(push_button_size.x/2+push_button_compliant_thumb_tolerence/2)]);
    translate([pi_board_size.x+2*pi_board_padding+walls_thickness,0,0])
    translate([0,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x-push_button_offset_from_corners,walls_thickness])
    translate([0,push_button_size.x,0])
    cube(size=[walls_thickness,push_button_compliant_thumb_tolerence,push_button_distance_from_case_bottom+push_button_size.x-(push_button_size.x/2+push_button_compliant_thumb_tolerence/2)]);
    
    translate([pi_board_size.x+2*pi_board_padding+walls_thickness,0,0])
    translate([0,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x-push_button_offset_from_corners,walls_thickness])
    translate([0,push_button_size.x/2+push_button_compliant_thumb_tolerence/2,push_button_distance_from_case_bottom+push_button_size.x-(push_button_size.x/2+push_button_compliant_thumb_tolerence/2)])
    rotate([90,0,90])
    linear_extrude(walls_thickness) arc(push_button_size.x/2-push_button_compliant_thumb_tolerence/2, [0,180], push_button_compliant_thumb_tolerence);
}

// Push button backplates
// 1
translate([walls_thickness,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x,walls_thickness])
translate([push_button_size.z,-push_button_offset_from_corners,push_button_distance_from_case_bottom])
cube(size=[push_button_support_thickness,push_button_size.x+push_button_offset_from_corners,push_button_size.y]);
// 2
translate([walls_thickness+pi_board_size.x+2*pi_board_padding,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x,walls_thickness])
translate([-push_button_size.z-push_button_support_thickness,-push_button_offset_from_corners,push_button_distance_from_case_bottom])
cube(size=[push_button_support_thickness,push_button_size.x+push_button_offset_from_corners,push_button_size.y]);

// Push button backplates supports (remove during post processing) 
// 1
translate([walls_thickness,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x,walls_thickness])
translate([push_button_size.z,-push_button_offset_from_corners,0])
translate([push_button_support_thickness/2,push_button_support_thickness/2,0])
cylinder(push_button_distance_from_case_bottom,push_button_support_thickness/2,push_button_support_thickness/2,$fn=resolution);
// 2
translate([walls_thickness+pi_board_size.x+2*pi_board_padding,pi_board_size.y+2*pi_board_padding+walls_thickness-push_button_size.x,walls_thickness])
translate([-push_button_size.z-push_button_support_thickness,-push_button_offset_from_corners,0])
translate([push_button_support_thickness/2,push_button_support_thickness/2,0])
cylinder(push_button_distance_from_case_bottom,push_button_support_thickness/2,push_button_support_thickness/2,$fn=resolution);