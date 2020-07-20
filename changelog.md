# Changelog
All notable changes to this project will be documented in this file.

## [1.0.1] - 2020-07-21
### Changed
- Allows rolls to be done with modifiers that are separated by blanks.  So 1d4 + 5 can work now where it previously would break.

## [1.0.0] - 2020-06-14
### Added
- Original implementation of all rolling mechanics
- Integration with a module that uses OS noise to generate random numbers (go-rint)
- Size limits on dice for the amount of dice rolled and the type of dice rolled
- Tests for specific inputs into the regex
