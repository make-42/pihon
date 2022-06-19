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
pi_board_hdmi_offset_from_edge = 46.85; //mm
pi_board_usb_size = [9,10,5]; //mm
pi_board_hdmi_size = [11.5,10,5]; //mm
pi_board_port_rounding_radius = 1; //mm

// Screens
screen_pcb_size = [24.7,27,13.3]; //mm (+- 0.2) 13.3mm are pin headers.
screen_pcb_screw_holes_center_from_board_edge_distance = 2.1; //mm
screen_pcb_screw_holes_diameter=2.8; //mm
screen_pcb_screw_holes_spacing=[20.5,23]; //mm
active_screen_area = [24.7,16.9,3.5]; //mm (old: 13.86)
active_screen_area_offset_from_corner = [0,5.08]; //mm

// Push buttons
push_button_size = [6,6,3.5]; //mm
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
cube(size=[pi_board_size.x+2*pi_board_padding+walls_thickness*2,pi_board_size.y+2*pi_board_padding+2*walls_thickness,walls_thickness]);
   // Fillet Edge
    fillet_mask_z(l=walls_thickness, r=rounding_radius, align=V_UP, $fn=resolution);
    
   translate([pi_board_size.x+2*pi_board_padding+walls_thickness*2,pi_board_size.y+2*pi_board_padding+2*walls_thickness,0])
    fillet_mask_z(l=walls_thickness, r=rounding_radius, align=V_UP, $fn=resolution);
    
   translate([pi_board_size.x+2*pi_board_padding+walls_thickness*2,0,0]) fillet_mask_z(l=walls_thickness, r=rounding_radius, align=V_UP, $fn=resolution);
    
   translate([0,pi_board_size.y+2*pi_board_padding+2*walls_thickness,0]) fillet_mask_z(l=walls_thickness, r=rounding_radius, align=V_UP, $fn=resolution);
    // Screen
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,-active_screen_area.y/2,0])
    cube([screen_pcb_size.x,active_screen_area.y,walls_thickness]);
    
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([0,-active_screen_area.y/2,0])
    cube([screen_pcb_size.x,active_screen_area.y,walls_thickness]);
    // Screws
   // Screen 1
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,active_screen_area.y/2+active_screen_area_offset_from_corner.y,0])
    translate([screen_pcb_screw_holes_center_from_board_edge_distance,-screen_pcb_screw_holes_center_from_board_edge_distance,0])
    cylinder(walls_thickness*2,screen_pcb_screw_holes_diameter/2,screen_pcb_screw_holes_diameter/2,$fn=resolution);
    
    translate([screen_pcb_screw_holes_spacing.x,0,0])
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,active_screen_area.y/2+active_screen_area_offset_from_corner.y,0])
    translate([screen_pcb_screw_holes_center_from_board_edge_distance,-screen_pcb_screw_holes_center_from_board_edge_distance,0])
    cylinder(walls_thickness*2,screen_pcb_screw_holes_diameter/2,screen_pcb_screw_holes_diameter/2,$fn=resolution);
    
    translate([0,-screen_pcb_screw_holes_spacing.y,0])
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,active_screen_area.y/2+active_screen_area_offset_from_corner.y,0])
    translate([screen_pcb_screw_holes_center_from_board_edge_distance,-screen_pcb_screw_holes_center_from_board_edge_distance,0])
    cylinder(walls_thickness*2,screen_pcb_screw_holes_diameter/2,screen_pcb_screw_holes_diameter/2,$fn=resolution);
    
    translate([screen_pcb_screw_holes_spacing.x,-screen_pcb_screw_holes_spacing.y,0])
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,active_screen_area.y/2+active_screen_area_offset_from_corner.y,0])
    translate([screen_pcb_screw_holes_center_from_board_edge_distance,-screen_pcb_screw_holes_center_from_board_edge_distance,0])
    cylinder(walls_thickness*2,screen_pcb_screw_holes_diameter/2,screen_pcb_screw_holes_diameter/2,$fn=resolution);
    // Screen 2
    translate([screen_pcb_size.x,0,0])
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,active_screen_area.y/2+active_screen_area_offset_from_corner.y,0])
    translate([screen_pcb_screw_holes_center_from_board_edge_distance,-screen_pcb_screw_holes_center_from_board_edge_distance,0])
    cylinder(walls_thickness*2,screen_pcb_screw_holes_diameter/2,screen_pcb_screw_holes_diameter/2,$fn=resolution);
    
    translate([screen_pcb_size.x,0,0])
    translate([screen_pcb_screw_holes_spacing.x,0,0])
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,active_screen_area.y/2+active_screen_area_offset_from_corner.y,0])
    translate([screen_pcb_screw_holes_center_from_board_edge_distance,-screen_pcb_screw_holes_center_from_board_edge_distance,0])
    cylinder(walls_thickness*2,screen_pcb_screw_holes_diameter/2,screen_pcb_screw_holes_diameter/2,$fn=resolution);
    
    translate([screen_pcb_size.x,0,0])
    translate([0,-screen_pcb_screw_holes_spacing.y,0])
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,active_screen_area.y/2+active_screen_area_offset_from_corner.y,0])
    translate([screen_pcb_screw_holes_center_from_board_edge_distance,-screen_pcb_screw_holes_center_from_board_edge_distance,0])
    cylinder(walls_thickness*2,screen_pcb_screw_holes_diameter/2,screen_pcb_screw_holes_diameter/2,$fn=resolution);
    
    translate([screen_pcb_size.x,0,0])
    translate([screen_pcb_screw_holes_spacing.x,-screen_pcb_screw_holes_spacing.y,0])
    translate([pi_board_size.x/2+pi_board_padding+walls_thickness,pi_board_size.y/2+pi_board_padding+walls_thickness,0])
    translate([-screen_pcb_size.x,active_screen_area.y/2+active_screen_area_offset_from_corner.y,0])
    translate([screen_pcb_screw_holes_center_from_board_edge_distance,-screen_pcb_screw_holes_center_from_board_edge_distance,0])
    cylinder(walls_thickness*2,screen_pcb_screw_holes_diameter/2,screen_pcb_screw_holes_diameter/2,$fn=resolution);
}