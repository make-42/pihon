# pihon

![pihon](https://raw.githubusercontent.com/make-42/pihon/master/Blender/output.gif)

A compact e-book reader.

## Freestyle renders

![pihon freestyle render 1](https://raw.githubusercontent.com/make-42/pihon/master/Blender/freestyle.png)
![pihon freestyle render 2](https://raw.githubusercontent.com/make-42/pihon/master/Blender/freestyle2.png)

## BOM (Prices in France as of June 2022)

- Raspberry Pi Zero WH (€ 10.44)
- 2x SSD1306 OLED Display (€ 1.73/each)
- 2x Momentary Switch (€ 0.50/each)
- Wires (€ 0.40)

Total: € 15.30 (excluding shipping)

# Raspberry Pi Zero W Boot Config

Don't forget to add this line to the `/boot/config.txt` file:

```
# Enable second i2c port
dtoverlay=i2c-gpio,bus=4,i2c_gpio_delay_us=1,i2c_gpio_sda=23,i2c_gpio_scl=24
```

## Wiring

![pihon Wiring](https://raw.githubusercontent.com/make-42/pihon/master/Schematic/Wiring.png)

## Schematic

![pihon Schematic](https://raw.githubusercontent.com/make-42/pihon/master/Schematic/Schematic.png)
