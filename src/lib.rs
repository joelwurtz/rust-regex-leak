use regex::Regex;

#[no_mangle]
pub extern "C" fn create_regex() -> *const Regex {
    let regex = Regex::new("(.+?)").expect("");

    Box::into_raw(Box::new(regex))
}

#[no_mangle]
pub unsafe extern "C" fn regex_match_leak(_regex: *const Regex) {
    let regex = &*_regex;

    regex.is_match("test");
}