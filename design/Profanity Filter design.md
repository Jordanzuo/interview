In order to filter profanity words, I introduce an algorithm wich is called DFA(Deterministic finite automaton) with the underlying data structure called Trie.

There are a few steps to implement this algorithm:
- Use a list of profanity words to construct the dfa object.
- Given a certain sentence, find if it contains any profanity word in the Trie.
- If found, then replaces those words to *.

See:
https://en.wikipedia.org/wiki/Deterministic_finite_automaton
https://en.wikipedia.org/wiki/Trie