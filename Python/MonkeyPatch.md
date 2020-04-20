# Monkey Patch

https://stackoverflow.com/questions/5626193/what-is-monkey-patching

No, it's not like any of those things. It's simply the dynamic replacement of attributes at runtime.

For instance, consider a class that has a method `get_data`. This method does an external lookup (on a database or web API, for example), and various other methods in the class call it. However, in a unit test, you don't want to depend on the external data source - so you dynamically replace the `get_data` method with a stub that returns some fixed data.

Because Python classes are mutable, and methods are just attributes of the class, you can do this as much as you like - and, in fact, you can even replace classes and functions in a module in exactly the same way.

But, as a [commenter](https://stackoverflow.com/users/2810305/lutz-prechelt) pointed out, use caution when monkeypatching:

1. If anything else besides your test logic calls `get_data` as well, it will also call your monkey-patched replacement rather than the original -- which can be good or bad. Just beware.
2. If some variable or attribute exists that also points to the `get_data` function by the time you replace it, this alias will not change its meaning and will continue to point to the original `get_data`. (Why? Python just rebinds the name `get_data` in your class to some other function object; other name bindings are not impacted at all.)

