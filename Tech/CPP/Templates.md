===template `class` or `typename`?===
http://blogs.msdn.com/b/slippman/archive/2004/08/11/212768.aspx
Summary: Stroustrup originally used class to specify types in templates to avoid introducing a new keyword. Some in the committee worried that this overloading of the keyword led to confusion. Later, the committee introduced a new keyword typename to resolve syntactic ambiguity, and decided to let it also be used to specify template types to reduce confusion, but for backward compatibility, class kept its overloaded meaning.
