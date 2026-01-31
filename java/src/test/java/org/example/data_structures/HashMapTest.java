package org.example.data_structures;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class HashMapTest {
    @Test
    public void testPutAndGet() {
        HashMap<String, Integer> map = new HashMap<>();
        map.put("one", 1);
        map.put("two", 2);
        assertEquals(1, map.get("one"));
        assertEquals(2, map.get("two"));
        map.put("one", 11); // Test duplicate key handling
        assertEquals(11, map.get("one"));
        map.put("key1", 3);
        map.put("key2", 5);
    }
}