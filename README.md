Vectormath for Go
=================

This is an adaptation of the scalar C functions from Sony's Vector Math
library, as found in the Bullet-2.79 source code. (Note that the current svn
for Bullet retains a subset of this library in a different location.) The full
library can be found in the following directory of the Bullet-2.79 source
tarball:

    Extras/vectormathlibrary/include/vectormath/scalar/c

Special thanks to Sony Computer Entertainment Inc. for allowing the use of
their Vector Math library under a BSD-style license, which made this possible.

Approach
--------

Start with a direct conversion of the Sony library to Go. (This was shipping
code, so we will assume it's correct.) Then iterate to make it more Go-like,
while trying to balance the twin goals of same/better performance and
maintaining readability.

The C version of the library, as opposed to the C++ version, was chosen as a
starting point mainly to avoid the distraction of function renaming, as the C++
version of Sony's library relies somewhat on operator and function overloads,
which have no direct equivalent in Go.

Note that the original C library provides two versions of each file, e.g.,
vec_aos.h and vec_aos_v.h. I've brought over the former as opposed to the
latter, since this simplifies things for beginners. The 'result' argument is
always a pointer passed as the first argument to a function. Currently, the
only types used in this library that are passed by value are simple types,
e.g., float and int. Everything else is a pointer. (This should, in theory,
simplify memory thrash considerations as well. In practice.. I have no idea. :)

Future Direction
----------------

Further research is required for determining:
- What makes the most sense for optimizing this code for SIMD
- Whether we should be passing vectors, etc., by value instead of reference
- If we should continue to declare vectors as individual values as opposed to
  an array of values. This would avoid the branching currently required to
  access members by index, but would disallow accessing members by common name
  (x, y, z and w) due to Golang's lack of a union equivalent.
- Whether it makes sense to stay with 32-bit, or if we would get the same
  performance with 64-bit

Feedback on this library is welcome and appreciated, though I make no promises
about my ability to deliver on anything beyond what you see here. :)

License
-------

I have licensed my modifications under a license similar to that of the
original library on which this is based. See the LICENSE file for more details.

