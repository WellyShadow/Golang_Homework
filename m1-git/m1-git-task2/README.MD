GIT Basics homework II
Part 1 (do 'GIT Basics homework' first)
1.Checkout on ”develop”.
2.Do all your experiments in folder “task2 –git pracrice II”.
3.Create any file innew folder “task2 –git pracrice II”and commit it.
4.Create newbranch “first”.
5.Create anothernew branch “second”.
6.Do one (or more) commits to “develop”(Add to repository some files, change them).
7.Checkout on the “first”branch and do several commits (about 5).
8.Checkout on the “second”branch and do several commits (about 3). Make sure that you have at least one commit, which contains changesinmore than one file.
9.Do interactive rebasing. Your commits from the “second”branch should appear on the top of “develop”branch. During interactive rebasing:
•Divide one of your commits (so you should have two commits instead of one);
•Change content of other commit and change commit message for it.
10.Merge “second” branch into develop(fast-forward).“Develop” and “second” branches should point to the same commit.
11.Checkout to the”first”branch.
12.Do interactive rebasing. Your commits from the “first”branch should appear on the top of “develop”branch. During interactive rebasing:
•Squash three commits into one;
•Drop one commit.13.Merge “first” branch into develop(fast-forward).Bothbranches should point to the same commit.
14.Merge the“develop”branchinto “master”branch.
15.Push all your local branches to origin.
16.Execute command “git reflog”and save it content somewhere (not in repository) with filename “GIT_Basics_homework II.txt”.
Additional task (*)

1)Create 4 commits into master branch with minor changes in file that is already exist.
2)Execute command “git reset--hard HEAD~4”
3)Are you ableto restore changes of third commit? Describe your steps how to do that.