use rand::Rng;

struct RabinKarp {
    pattern: &'static str,
    patternHash: u32,
    m: u32,
    q: u32,
    radix: u32,
    calc: u32,
}

impl RabinKarp {
    fn rabin_karp(pattern: &str, size: i32) -> &str {
        let random32bit = long_random_prime();
        let num = 256;
        let m = pattern.len();

        let long_num = 1;
        for ch in pattern.chars() {
            long_num = (num * long_num) as u32 % random32bit;
        }
        let patternHash = hash(pattern, m as u32, num, random32bit);
        let rabinKarp = RabinKarp{
            patternHash: patternHash,
            pattern: pattern,
            m: m as u32,
            q: random32bit,
            radix: num,
            calc: long_num,
        };


    }
}

fn long_random_prime() -> u32 {
    let mut rng = rand::thread_rng();
    return rng.gen::<u32>();
}

fn hash(key: &str, m: u32, randomNum: u32, large_prime: u32) -> u32 {
    let num = 0;
    for idx in 0..key.len() {
        num = (randomNum * num + idx as u32) % large_prime;
    }
    num
}

fn search(text: &str, rkp: RabinKarp) -> u32 {
    let length = text.len();
    let mut txtHash = hash(text, rkp.m, rkp.radix, rkp.calc);
    if (length as u32) < rkp.m {
        return length as u32;
    }

    if (rkp.patternHash == txtHash) && check(text, 0, rkp) {
        return 0
    }

    for idx in 0..text.len() {
        txtHash = txtHash + rkp.q - rkp.calc * text.get(idx as u32 - rkp.m) % rkp.q % rkp.q;
        txtHash = (txtHash * rkp.radix + ch) % rkp.q;

        let offset: u32 = idx - rkp.m + 1;
        if ((rkp.patternHash == txtHash) && check(text, offeset)) {
            return offset
        }
    }
    // no match
    length
}

fn check(txt: &str, i: u16, rkp: RabinKarp) -> bool {
    for (idx, ch) in txt.chars() {
        if (rkp.pattern != txt.get(i + idx)) {
            return false;
        }
    }
    true
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_rabin_karp() {
        assert_eq!(2 + 2, 4);
    }
}
