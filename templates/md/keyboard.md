# Introduction
<!-- `{% func KeyboardIntroduction() %}` -->


If you are typing a couple (or more) hours a day, you should consider investing in a keyboard that would make typing more comfortable. It can be beneficial in the long run.

That's why right after I started my first job I decided to buy a mechanical keyboard: REDRAGON Fizz K617.

<br>


## REDRAGON Fizz K617

It is a 60% keyboard, meaning that it only has around 60% of the keys of a normal full keyboard: you need to press a combination of keys to access the missing keys.

I changed the default switches into some silent light ones (Outemu silent white) so that I would feel less fatigue and I wouldn't make much noise from typing.



## Colemak

Since I wanted to make typing even more ergonomic, I also decided to learn to type with a new layout: Colemak.

It is one of the layouts optimized for touch typing in English: the most frequently used letters are moved to the home row to make them easier to press.

<br>

**![Colemak](./assets/keyboard/colemak.png)**<!-- class="mx-auto h-full" -->

<br>

Compared to the Qwerty layout, the positions of 17 of keys are changed but common shortcuts (like copy-pasta) are kept the same.

<br>

[**https://colemak.com/**](https://colemak.com/)<!-- class="animate-pulse bg-gray-200" -->

<br>

<!-- `{% endfunc %}` -->


# Split
<!-- `{% func KeyboardSplit() %}` -->

After a year and a half of using the 60% keyboard, I looked into some **custom split ergonomical keyboards**<!-- class="bg-orange-200" -->.

There are some issues (kind of) that split keyboard can improve.

<br>

## Arms position

To type on a normal keyboard you have to bend your arms toward the center of your desk where the keyboard is located.

With a split keyboard, your hands stay at shoulder distance (or at any distance of your preference) with **less strain on your shoulder and arms**<!-- class="bg-green-200" -->.

<br>


## Keys position

**In a normal keyboard, the keys are staggered**<!-- class="bg-red-200" --> and in order to touch type you should assign a set of keys to each of your fingers.

**![Qwerty Layout](./assets/keyboard/layout.png)**<!-- class="mx-auto h-full" -->

<br>

Try to type the word 'byzantine'. Did you find it comfortable to type the first three letters?

<br>

Because of the staggered layout, B and Y are the furthest keys assigned to index fingers (relative to F and J which are the base positions). I often misstype their Colemak counterparts because of it.

<br>

According to many touch typing guides, letter Z should be pressed with the left pinkie.

Both hands are already bent toward the center, why strain the outer and weakest finger even more to type Z? I find it more comfortable to type it with my left ring finger.

<br>

**Split keyboards can be ortholinear**<!-- class="bg-indigo-200" -->, solving these problems.

<br>

## Less keys and less finger movements

Some split keyboard are 40%. The idea is that there are multiple layers that can be activated by holding some specific keys.

This means that you need to press multiple keys together to type, for example, symbols or to use the F keys (F1, F2, F3 etc).

<br>

The advantage is that **the fingers don't move that much to type**<!-- class="bg-teal-200" -->: it is easier (and even faster once you memorize the combo) to press a coupe of keys in easy to reach positions than to press a single key that requires to lift and move the entire hand (some example: backspace, enter and arrows).

<!-- `{% endfunc %}` -->


# Circuit
<!-- `{% func KeyboardCircuit() %}` -->

After learning how these custom split keyboards works **I decided to design and to build one myself**<!-- class="bg-teal-200" --> and I used it to type all of this.

<br>

There are two types of custom keyboard based on how the circuit is made: hand wired and PCB (printed circuit board).

I went for the latter as it looked easier.

<br>

The tools I used to generate the PCB design are [**ergogen**](https://github.com/ergogen/ergogen)<!-- class="bg-purple-200 animate-pulse" --> (to draw the layout) and [**KiCad**](https://www.kicad.org/)<!-- class="bg-orange-200 animate-pulse" --> (to generate the electrical connections and the final production files) and the basic design was **~~stolen~~** inspired by another custom keyboard:
[**cheapino**](https://github.com/tompi/cheapino)<!-- class="bg-yellow-200 animate-pulse" -->

<br>

I also added two rotary encoders in the design but after the assembling they didn't work as I wanted it to and I now feel like they aren't that necessary to have.

<br>


## Design

Here is my final circuit design.

It can be used **for both left and right splits**<!-- class="bg-green-200" --> and **the microcontroller can be installed on either side**<!-- class="bg-teal-200" -->.

**30 keys in total**<!-- class="bg-blue-200" -->, 15 on each side, 3 keys assigned to each finger.

<br>

**![Circuit](./assets/keyboard/kicad.png)**<!-- class="mx-auto h-full" -->

<br>


## How it works

On each key there is a switch that, once pressed, connects the two side electrically and closes the circuit letting the current pass through.

<br>

Each side of the switch is connected to a pin of the MCU ([**RP2040 Zero**](https://www.waveshare.com/wiki/RP2040-Zero)<!-- class="bg-gray-200 animate-pulse" --> in my case) that elaborates what needs to happen when a specific key is pressed.

<br>

Usually the microcontroller has a very limited amount of pins, less than the number of keyboard keys: in order to make it work a diode is connected to each key and a matrix arrangement of the connections is used.

The standard configuration is to have one pin of the MCU for each of the rows and columns of the keyboard.

<br>

On each cycle the MCU scans the pins, checking if there are other pins with a signal. The scanning direction is set to 'row to column' or 'column to row'.

Each switch is connected to a distinct (row pin, column pin) pair so that it can be identified when pressed.

<br>

If (row1, col1) and (row2, col2) keys are pressed at the same time then it is possible for the MCU to also register (row1, col2) or (row2, col1): a diode on each key allows the current to flow only on one direction and prevents this effect called **ghosting**.

<br>

Given N number of free MCU pins, the maximum number of keyboard keys that can be registered with a regular matrix scanning method is (N/2)*(N/2) using (N/2) row pins and (N/2) column pins.

<br>

This number can be doubled using a [**duplex matrix**](https://kbd.news/The-Japanese-duplex-matrix-1391.html)<!-- class="bg-pink-200 animate-pulse" -->: the scanning goes both ways and for each (row, col) pair there are 2 switches.

The difference between these 2 switches is the orientation of the diodes (row2col or col2row) so that the scanning direction can also be used to identify the pressed key.

<br>


An [**RJ45 cable**](https://en.wikipedia.org/wiki/Modular_connector#8P8C)<!-- class="bg-amber-200 animate-pulse" --> connects the two side: this allows for 8 connection pins to be used.



## JLCPCB

After finishing the PCB design I needed to have it printed so I uploaded the gerber file to [**JLCPCB**](https://jlcpcb.com/)<!-- class="bg-cyan-200 animate-pulse" -->.

<br>

I really like their service because they offer a special deal if the PCB is less than 10cm X 10cm in size and I designed the circuit within these constraints.

The total cost of 5 boards + shipping (to Italy) + taxes was only around 4€.


## Aliexpress

The rest of the materials (soldering tools, diodes, switches, keycaps, MCU, cables etc.) were bought from [**Aliexpress**](https://aliexpress.com/)<!-- class="bg-red-200 animate-pulse" -->

On Aliexpress there are actually some good quality products at very low prices that would be nearly impossible to find and buy locally.


<!-- `{% endfunc %}` -->

# Assembling
<!-- `{% func KeyboardAssembling() %}` -->

Here some pictures of the assembling process.

<br>

**![](./assets/keyboard/a1.jpg)**<!-- class="mx-auto h-full" -->

**![](./assets/keyboard/a2.jpg)**<!-- class="mx-auto h-full" -->

**![](./assets/keyboard/a3.jpg)**<!-- class="mx-auto h-full" -->

**![](./assets/keyboard/a4.jpg)**<!-- class="mx-auto h-full" -->

**![](./assets/keyboard/a5.jpg)**<!-- class="mx-auto h-full" -->

**![](./assets/keyboard/a6.jpg)**<!-- class="mx-auto h-full" -->

**![](./assets/keyboard/a7.jpg)**<!-- class="mx-auto h-full" -->

<br>

In the last picture if you are wondering what the extra small little cable is for, well... let's say that there was a ***Happy Little Accident*** during the design of the PCB.

<br>

An RJ45 PCB connection has 8 pins and they are positioned in 2 staggered rows, as if they were forming a parallelogram.

<br>

Since it is not symmetric, **I added an extra pin hole**<!-- class="bg-teal-200" --> so that the board could be flipped and be used both as left and right split **but I connected the pinholes to the MCU as if there were 9 available connecting pins while only 8 are effectively in use**<!-- class="bg-red-200" -->.

<!-- `{% endfunc %}` -->


# Firmware
<!-- `{% func KeyboardFirmware() %}` -->

The last step is to install a firmware so that the MCU can send key presses when a switch is pushed.

The most widely used firmware for custom keyboard is [**QMK**](https://github.com/qmk/qmk_firmware)<!-- class="bg-indigo-200 animate-pulse" --> (based on C)...

<br>

... but I decided to go with [**PRK**](https://github.com/picoruby/prk_firmware/wiki/Tutorial)<!-- class="bg-red-200 animate-pulse" --> (based on Ruby).

<br>

That's the reason I put an image of Poro (from LOL) on the MCU since PORO can also stand for 'Plain Old Ruby Object'

<br>

The PRK firmware is very easy to install following the documentation and all that is needed after that is to upload a `keymap.rb` file that contains the definition of the pins, the matrix and the assigment of the keys.

<br>


## Keymap

The layout I'm using is basically Colemak but with some modifications.

<br>

I am missing the central 6 letter keys assigned to the index fingers which in Colemak are G, D, B, J, H and K (in Qwerty: T, G, B, Y, H, N).

<br>

G, D, B and H have replaced comma, dot, slash and semicolons of the standard Colemak, placed on the right side of the keyboard.

<br>

J, K, all the numbers, symbols, F keys and navigations have been placed in secondary layers.

<br>

In the thumb cluster I have Ctrl, Space, Shift, Backspace and two keys to activate the secondary layers.

<br>

Overall I can feel some improvements in my posture when typing.

<!-- `{% endfunc %}` -->