#!/usr/bin/env bash

EXP="// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file."

RC=0

for f in $(find ./ -name '*.go'); do


GOT=$(head -n 3 $f)
if [ ! "$(head -n 3 $f)" = "$EXP" ] ; then
  echo "$f does not have a (valid) copyright header"
  RC=1
fi

done

exit $RC

