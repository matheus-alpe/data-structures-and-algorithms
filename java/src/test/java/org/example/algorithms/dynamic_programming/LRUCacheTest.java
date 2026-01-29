package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class LRUCacheTest {
    @Test
    public void testLRUCache() {
        LRUCache cache = new LRUCache(2);
        cache.put(1, 1);
        cache.put(2, 2);
        assertEquals(1, cache.get(1));
        cache.put(3, 3);
        assertEquals(-1, cache.get(2));
        cache.put(4, 4);
        assertEquals(-1, cache.get(1));
        assertEquals(3, cache.get(3));
        assertEquals(4, cache.get(4));
    }

    @Test
    public void testLRUCacheUpdate() {
        LRUCache cache = new LRUCache(2);
        cache.put(1, 1);
        cache.put(2, 2);
        cache.put(1, 10); // update value for key 1
        assertEquals(10, cache.get(1));
        cache.put(3, 3);
        assertEquals(-1, cache.get(2)); // key 2 should be evicted
    }

    @Test
    public void testLRUCacheCapacityOne() {
        LRUCache cache = new LRUCache(1);
        cache.put(1, 1);
        assertEquals(1, cache.get(1));
        cache.put(2, 2);
        assertEquals(-1, cache.get(1)); // key 1 should be evicted
        assertEquals(2, cache.get(2));
    }

    @Test
    public void testLRUCacheGetNonExistentKey() {
        LRUCache cache = new LRUCache(2);
        assertEquals(-1, cache.get(5)); // key 5 does not exist
        cache.put(1, 1);
        assertEquals(-1, cache.get(2)); // key 2 does not exist
    }
}