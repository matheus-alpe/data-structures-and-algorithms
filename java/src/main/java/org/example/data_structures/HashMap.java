package org.example.data_structures;

public class HashMap<T, V> {
    private final int size;
    private final Entry[] entries;

    @SuppressWarnings("unchecked")
    public HashMap(int size) {
        this.size = size;
        this.entries = (Entry[]) new HashMap<?, ?>.Entry[size];
    }

    public HashMap() {
        this(10);
    }

    public void put(T key, V value) {
        int index = hash(key);
        Entry newEntry = new Entry(key, value);
        Entry current = entries[index];

        if (current == null) {
            entries[index] = newEntry;
        } else {
            while (current.next != null) {
                if (current.key.equals(key)) {
                    current.value = value;
                    return;
                }
                current = current.next;
            }
            if (current.key.equals(key)) {
                current.value = value;
                return;
            }
            current.next = newEntry;
        }
    }

    public V get(T key) {
        int index = hash(key);
        Entry current = entries[index];

        while (current != null) {
            if (current.key.equals(key)) {
                return current.value;
            }
            current = current.next;
        }

        return null;
    }

    private int hash(T key) {
        return (key.hashCode() & Integer.MAX_VALUE) % size;
    }

    private class Entry {
        final T key;
        V value;
        Entry next;

        Entry(T key, V value) {
            this.key = key;
            this.value = value;
        }
    }
}
