# Remove N-th Node from End of List

> Striver SDE Sheet 028 · Day 5 Linked List — two pointers a fixed n apart; when the lead one hits the end, the trailer sits before the victim. Go, stdlib only, tested.

## Problem

Delete the n-th node from the *end* of a singly linked list in a single pass — without first measuring the length.

## Approach

Send a `fast` pointer n steps ahead, then advance `fast` and `slow` together until `fast` reaches the last node. The gap is constant, so `slow` lands exactly one node before the one to delete — which is the position you need, since singly linked deletion rewires the predecessor. Starting both pointers at a dummy node makes "delete the head" (n equals the length) the same code path as any other deletion.

**Complexity:** O(n) time, O(1) space, single pass.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
