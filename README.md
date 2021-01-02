# ThompsonCalculator

This project will build and maintain a gui in GoLang for a Thompson Groups Calculator.

The project has as it main calculating back-end engine the project github.com/loeksnokes/TreePair, which implements a pair of prefix codes on an arbitrary alphabet size, together with a bijection between them (thus, representing an element of V_n for whatever value of n).  TreePair itself of course depends on github.com/loeksnokes/PrefCode, which implements a prefix code on a fixed alphabet as given by user.  These back-end pieces are still under construction (e.g., currently (02/01/2021) error-handling is weak) but they currently do support most standard activities including multiplying and reducing elements, and verifying if a given element is in F_n, T_n, or V_n.

For the front-end, we are in the process of implementing the gui interface using the pixel 2D programming library at github.com/faiface/pixel.  There will also be a nlp parser window to interact with the back-end (e.g., build, invert, multiply, reduce elements), and pop-up windows for visualisation of implemented elements.  It is intended that there will also be a commutator calculator using Thurston Diagrams, and if my interest remains strong, eventually there will be pop-ups for revealing pairs or their various generalisations as described in Ewa Bieniecka's PhD Thesis.
