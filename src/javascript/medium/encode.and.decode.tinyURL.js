'use strict';

/**
 * TinyURL is a URL shortening service where you enter a URL such as https://leetcode.com/problems/design-tinyurl 
 * and it returns a short URL such as http://tinyurl.com/4e9iAk.
 * Design the encode and decode methods for the TinyURL service. 
 * There is no restriction on how your encode/decode algorithm should work. 
 * You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.
 */

const map = {};

/**
 * Encodes a URL to a shortened URL.
 *
 * @param {string} longUrl
 * @return {string}
 */
const encode = function(longUrl) {
    let shortUrl = 'http://tinyurl.com/' + map.size;
    // map[longUrl] = shortUrl;
    map[shortUrl] = longUrl;
    return shortUrl;
};

/**
 * Decodes a shortened URL to its original URL.
 *
 * @param {string} shortUrl
 * @return {string}
 */
const decode = function(shortUrl) {
    return map[shortUrl];
};

const longUrl = 'http://baidu.com/user/info/list/34245322244';
const shortUrl = encode(longUrl);
console.log(shortUrl);
console.log(decode(shortUrl));
